/*******************************************************************************
 * Copyright 2019 Dell Inc.
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
 *
 *******************************************************************************/

package proxy

import (
	"context"
	"fmt"
	"os"
	"sync"

	bootstrapContainer "github.com/edgexfoundry/edgex-go/internal/pkg/bootstrap/container"
	"github.com/edgexfoundry/edgex-go/internal/pkg/bootstrap/startup"
	"github.com/edgexfoundry/edgex-go/internal/pkg/di"
	"github.com/edgexfoundry/edgex-go/internal/security/proxy/config"

	"github.com/edgexfoundry/go-mod-core-contracts/clients/logger"
)

// Global variables
var Configuration = &config.ConfigurationStruct{}
var LoggingClient logger.LoggingClient

type Bootstrap struct {
	insecureSkipVerify bool
	initNeeded         bool
	resetNeeded        bool
	userTobeCreated    string
	userOfGroup        string
	userToBeDeleted    string
}

func NewBootstrapHandler(
	insecureSkipVerify bool,
	initNeeded bool,
	resetNeeded bool,
	userTobeCreated string,
	userOfGroup string,
	userToBeDeleted string) *Bootstrap {

	return &Bootstrap{
		insecureSkipVerify: insecureSkipVerify,
		initNeeded:         initNeeded,
		resetNeeded:        resetNeeded,
		userTobeCreated:    userTobeCreated,
		userOfGroup:        userOfGroup,
		userToBeDeleted:    userToBeDeleted,
	}
}

func (b *Bootstrap) errorAndHalt(message string) {
	LoggingClient.Error(message)
	os.Exit(1)
}

func (b *Bootstrap) haltIfError(err error) {
	if err != nil {
		b.errorAndHalt(err.Error())
	}
}

// BootstrapHandler fulfills the BootstrapHandler contract and performs initialization needed by the data service.
func (b *Bootstrap) Handler(wg *sync.WaitGroup, ctx context.Context, startupTimer startup.Timer, dic *di.Container) bool {
	// update global variables.
	LoggingClient = bootstrapContainer.LoggingClientFrom(dic.Get)

	req := NewRequestor(b.insecureSkipVerify, Configuration.Writable.RequestTimeout)
	if req == nil {
		os.Exit(1)
	}

	s := NewService(req)
	b.haltIfError(s.CheckProxyServiceStatus())

	if b.initNeeded {
		if b.resetNeeded {
			b.errorAndHalt("can't run initialization and reset at the same time for security service")
		}
		cert := NewCertificateLoader(req, Configuration.SecretService.CertPath, Configuration.SecretService.TokenPath)
		b.haltIfError(s.Init(cert)) // Where the Service init is called
	} else if b.resetNeeded {
		b.haltIfError(s.ResetProxy())
	}

	if b.userTobeCreated != "" && b.userOfGroup != "" {
		c := NewConsumer(b.userTobeCreated, req)
		b.haltIfError(c.Create(EdgeXKong))
		b.haltIfError(c.AssociateWithGroup(b.userOfGroup))

		t, err := c.CreateToken()
		if err != nil {
			b.errorAndHalt(fmt.Sprintf("failed to create access token for edgex service due to error %s", err.Error()))
		}

		fmt.Println(fmt.Sprintf("the access token for user %s is: %s. Please keep the token for accessing edgex services", b.userTobeCreated, t))

		file, err := os.Create(Configuration.KongAuth.OutputPath)
		b.haltIfError(err)

		utp := &UserTokenPair{User: b.userTobeCreated, Token: t}
		b.haltIfError(utp.Save(file))
	}

	if b.userToBeDeleted != "" {
		t := NewConsumer(b.userToBeDeleted, req)
		b.haltIfError(t.Delete())
	}

	return false
}
