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

package data

import (
	"context"
	"sync"

	dataContainer "github.com/edgexfoundry/edgex-go/internal/core/data/container"
	"github.com/edgexfoundry/edgex-go/internal/core/data/v2"
	errorContainer "github.com/edgexfoundry/edgex-go/internal/pkg/container"
	"github.com/edgexfoundry/edgex-go/internal/pkg/errorconcept"

	"github.com/edgexfoundry/go-mod-bootstrap/v2/bootstrap/container"
	"github.com/edgexfoundry/go-mod-bootstrap/v2/bootstrap/startup"
	"github.com/edgexfoundry/go-mod-bootstrap/v2/di"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/clients"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/clients/metadata"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/clients/urlclient/local"

	"github.com/gorilla/mux"
)

// Bootstrap contains references to dependencies required by the BootstrapHandler.
type Bootstrap struct {
	router *mux.Router
}

// NewBootstrap is a factory method that returns an initialized Bootstrap receiver struct.
func NewBootstrap(router *mux.Router) *Bootstrap {
	return &Bootstrap{
		router: router,
	}
}

// BootstrapHandler fulfills the BootstrapHandler contract and performs initialization needed by the data service.
func (b *Bootstrap) BootstrapHandler(ctx context.Context, wg *sync.WaitGroup, startupTimer startup.Timer, dic *di.Container) bool {
	loadRestRoutes(b.router, dic)
	v2.LoadRestRoutes(b.router, dic)

	configuration := dataContainer.ConfigurationFrom(dic.Get)
	lc := container.LoggingClientFrom(dic.Get)

	mdc := metadata.NewDeviceClient(local.New(configuration.Clients[clients.CoreMetaDataServiceKey].Url() + clients.ApiDeviceRoute))
	msc := metadata.NewDeviceServiceClient(local.New(configuration.Clients[clients.CoreMetaDataServiceKey].Url() + clients.ApiDeviceRoute))

	chEvents := make(chan interface{}, 100)
	// initialize event handlers
	initEventHandlers(lc, chEvents, mdc, msc, configuration)

	dic.Update(di.ServiceConstructorMap{
		dataContainer.MetadataDeviceClientName: func(get di.Get) interface{} {
			return mdc
		},
		dataContainer.EventsChannelName: func(get di.Get) interface{} {
			return chEvents
		},
		errorContainer.ErrorHandlerName: func(get di.Get) interface{} {
			return errorconcept.NewErrorHandler(lc)
		},
	})

	return true
}
