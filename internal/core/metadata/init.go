/*******************************************************************************
 * Copyright 2017 Dell Inc.
 * Copyright (c) 2019 Intel Corporation
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing permissions and limitations under
 * the License.
 *******************************************************************************/
package metadata

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/edgexfoundry/go-mod-core-contracts/clients"
	"github.com/edgexfoundry/go-mod-core-contracts/clients/coredata"
	"github.com/edgexfoundry/go-mod-core-contracts/clients/logger"
	"github.com/edgexfoundry/go-mod-core-contracts/clients/notifications"
	"github.com/edgexfoundry/go-mod-core-contracts/clients/types"
	registryTypes "github.com/edgexfoundry/go-mod-registry/pkg/types"
	"github.com/edgexfoundry/go-mod-registry/registry"

	"github.com/edgexfoundry/edgex-go/internal"
	"github.com/edgexfoundry/edgex-go/internal/core/metadata/interfaces"
	"github.com/edgexfoundry/edgex-go/internal/pkg/config"
	"github.com/edgexfoundry/edgex-go/internal/pkg/db"
	"github.com/edgexfoundry/edgex-go/internal/pkg/db/mongo"
	"github.com/edgexfoundry/edgex-go/internal/pkg/db/redis"
	"github.com/edgexfoundry/edgex-go/internal/pkg/startup"
	"github.com/edgexfoundry/edgex-go/internal/pkg/telemetry"
)

// Global variables
var Configuration *ConfigurationStruct
var dbClient interfaces.DBClient
var LoggingClient logger.LoggingClient
var registryClient registry.Client
var nc notifications.NotificationsClient
var vdc coredata.ValueDescriptorClient
var registryErrors chan error        //A channel for "config wait errors" sourced from Registry
var registryUpdates chan interface{} //A channel for "config updates" sourced from Registry.

func Retry(useRegistry bool, useProfile string, timeout int, wait *sync.WaitGroup, ch chan error) {
	until := time.Now().Add(time.Millisecond * time.Duration(timeout))
	for time.Now().Before(until) {
		var err error
		//When looping, only handle configuration if it hasn't already been set.
		if Configuration == nil {
			Configuration, err = initializeConfiguration(useRegistry, useProfile)
			if err != nil {
				ch <- err
				if !useRegistry {
					//Error occurred when attempting to read from local filesystem. Fail fast.
					close(ch)
					wait.Done()
					return
				}
			} else {
				// Initialize notificationsClient based on configuration
				initializeClients(useRegistry)
				// Setup Logging
				logTarget := setLoggingTarget()
				LoggingClient = logger.NewClient(clients.CoreMetaDataServiceKey, Configuration.Logging.EnableRemote, logTarget, Configuration.Writable.LogLevel)
			}
		}

		//Only attempt to connect to database if configuration has been populated
		if Configuration != nil {
			err := connectToDatabase()
			if err != nil {
				ch <- err
			} else {
				break
			}
		}
		time.Sleep(time.Second * time.Duration(1))
	}
	close(ch)
	wait.Done()

	return
}

func Init(useRegistry bool) bool {
	if Configuration == nil || dbClient == nil {
		return false
	}
	if useRegistry {
		registryErrors = make(chan error)
		registryUpdates = make(chan interface{})
		go listenForConfigChanges()
	}

	go telemetry.StartCpuUsageAverage()

	return true
}

func Destruct() {
	if dbClient != nil {
		dbClient.CloseSession()
		dbClient = nil
	}
}

func initializeConfiguration(useRegistry bool, useProfile string) (*ConfigurationStruct, error) {
	//We currently have to load configuration from filesystem first in order to obtain RegistryHost/Port
	configuration := &ConfigurationStruct{}
	err := config.LoadFromFile(useProfile, configuration)
	if err != nil {
		return nil, err
	}
	configuration.Registry = config.OverrideFromEnvironment(configuration.Registry)

	if useRegistry {
		err = connectToRegistry(configuration)
		if err != nil {
			return nil, err
		}

		rawConfig, err := registryClient.GetConfiguration(configuration)
		if err != nil {
			return nil, fmt.Errorf("could not get configuration from Registry: %v", err.Error())
		}

		actual, ok := rawConfig.(*ConfigurationStruct)
		if !ok {
			return nil, fmt.Errorf("configuration from Registry failed type check")
		}

		configuration = actual

		// Check that information was successfully read from Registry
		if configuration.Service.Port == 0 {
			return nil, errors.New("error reading configuration from Registry")
		}
	}
	return configuration, nil
}

func connectToRegistry(conf *ConfigurationStruct) error {
	var err error
	registryConfig := registryTypes.Config{
		Host:            conf.Registry.Host,
		Port:            conf.Registry.Port,
		Type:            conf.Registry.Type,
		ServiceKey:      clients.CoreMetaDataServiceKey,
		ServiceHost:     conf.Service.Host,
		ServicePort:     conf.Service.Port,
		ServiceProtocol: conf.Service.Protocol,
		CheckInterval:   conf.Service.CheckInterval,
		CheckRoute:      clients.ApiPingRoute,
		Stem:            internal.ConfigRegistryStemCore + internal.ConfigMajorVersion,
	}

	registryClient, err = registry.NewRegistryClient(registryConfig)
	if err != nil {
		return fmt.Errorf("connection to Registry could not be made: %v", err.Error())
	}

	// Check if registry service is running
	if !registryClient.IsAlive() {
		return fmt.Errorf("registry is not available")
	}

	// Register the service with Registry
	err = registryClient.Register()
	if err != nil {
		return fmt.Errorf("could not register service with Registry: %v", err.Error())
	}

	return nil
}

func listenForConfigChanges() {
	if registryClient == nil {
		LoggingClient.Error("listenForConfigChanges() registry client not set")
		return
	}

	registryClient.WatchForChanges(registryUpdates, registryErrors, &WritableInfo{}, internal.WritableKey)

	signals := make(chan os.Signal)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)

	for {
		select {
		case <-signals:
			// Quietly and gracefully stop when SIGINT/SIGTERM received
			return

		case ex := <-registryErrors:
			LoggingClient.Error(ex.Error())
		case raw, ok := <-registryUpdates:
			if ok {
				actual, ok := raw.(*WritableInfo)
				if !ok {
					LoggingClient.Error("listenForConfigChanges() type check failed")
				}

				Configuration.Writable = *actual

				LoggingClient.Info("Writeable configuration has been updated from the Registry")
				LoggingClient.SetLogLevel(Configuration.Writable.LogLevel)
			} else {
				return
			}
		}
	}
}

func connectToDatabase() error {
	var err error

	dbClient, err = newDBClient(Configuration.Databases["Primary"].Type)
	if err != nil {
		dbClient = nil
		return fmt.Errorf("couldn't create database client: %v", err.Error())
	}

	return nil
}

// Return the dbClient interface
func newDBClient(dbType string) (interfaces.DBClient, error) {
	switch dbType {
	case db.MongoDB:
		dbConfig := db.Configuration{
			Host:         Configuration.Databases["Primary"].Host,
			Port:         Configuration.Databases["Primary"].Port,
			Timeout:      Configuration.Databases["Primary"].Timeout,
			DatabaseName: Configuration.Databases["Primary"].Name,
			Username:     Configuration.Databases["Primary"].Username,
			Password:     Configuration.Databases["Primary"].Password,
		}
		return mongo.NewClient(dbConfig)
	case db.RedisDB:
		dbConfig := db.Configuration{
			Host: Configuration.Databases["Primary"].Host,
			Port: Configuration.Databases["Primary"].Port,
		}
		return redis.NewClient(dbConfig, LoggingClient) //TODO: Verify this also connects to Redis
	default:
		return nil, db.ErrUnsupportedDatabase
	}
}

func initializeClients(useRegistry bool) {
	// Create notification client
	nParams := types.EndpointParams{
		ServiceKey:  clients.SupportNotificationsServiceKey,
		Path:        clients.ApiNotificationRoute,
		UseRegistry: useRegistry,
		Url:         Configuration.Clients["Notifications"].Url() + clients.ApiNotificationRoute,
		Interval:    Configuration.Service.ClientMonitor,
	}

	nc = notifications.NewNotificationsClient(nParams, startup.Endpoint{RegistryClient: &registryClient})

	vParams := types.EndpointParams{
		ServiceKey:  clients.CoreDataServiceKey,
		Path:        clients.ApiValueDescriptorRoute,
		UseRegistry: useRegistry,
		Url:         Configuration.Clients["CoreData"].Url() + clients.ApiValueDescriptorRoute,
		Interval:    Configuration.Service.ClientMonitor,
	}
	vdc = coredata.NewValueDescriptorClient(vParams, startup.Endpoint{RegistryClient: &registryClient})
}

func setLoggingTarget() string {
	if Configuration.Logging.EnableRemote {
		return Configuration.Clients["Logging"].Url() + clients.ApiLoggingRoute
	}
	return Configuration.Logging.File
}
