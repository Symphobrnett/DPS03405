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
 * @author: Daniel Harms, Dell
 *******************************************************************************/

package secretstore

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/edgexfoundry/go-mod-core-contracts/clients/logger"
)

func NewRequester(skipVerify bool, loggingClient logger.LoggingClient) *http.Client {
	var tr *http.Transport

	if skipVerify {
		tr = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: skipVerify},
		}
	} else {
		caCert, err := ioutil.ReadFile(Configuration.SecretService.CaFilePath)
		if err != nil {
			loggingClient.Error("failed to load rootCA certificate.")
			return nil
		}
		loggingClient.Info("successfully loaded rootCA certificate.")
		caCertPool := x509.NewCertPool()
		caCertPool.AppendCertsFromPEM(caCert)

		tr = &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs:            caCertPool,
				InsecureSkipVerify: skipVerify,
			},
			TLSHandshakeTimeout: 10 * time.Second,
		}
	}
	return &http.Client{Timeout: 10 * time.Second, Transport: tr}
}
