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
 *******************************************************************************/

package certificates

import (
	"github.com/edgexfoundry/go-mod-core-contracts/clients/logger"
	"testing"

	"github.com/edgexfoundry/edgex-go/internal/security/setup"
)

func TestGeneratePrivateKey(t *testing.T) {
	baseSeed := setup.CertificateSeed{
		RSAScheme:  false,
		ECScheme:   false,
		RSAKeySize: setup.RSA_4096,
		ECCurve:    setup.EC_224,
	}

	rsaSeed := baseSeed
	rsaSeed.RSAScheme = true

	ecSeed := baseSeed
	ecSeed.ECScheme = true

	ec256Seed := ecSeed
	ec256Seed.ECCurve = setup.EC_256

	ec384Seed := ecSeed
	ec384Seed.ECCurve = setup.EC_384

	ec521Seed := ecSeed
	ec521Seed.ECCurve = setup.EC_521

	logger := logger.MockLogger{}

	tests := []struct {
		name        string
		seed        setup.CertificateSeed
		expectError bool
	}{
		{"BaseFail", baseSeed, true},
		{"RSAPass", rsaSeed, false},
		{"ECPass", ecSeed, false},
		{"EC256Pass", ec256Seed, false},
		{"EC384Pass", ec384Seed, false},
		{"EC521Pass", ec521Seed, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := generatePrivateKey(tt.seed, logger)
			if err != nil && !tt.expectError {
				t.Error(err)
				return
			}
			if err == nil && tt.expectError {
				t.Error("expected error but none was thrown")
				return
			}
		})
	}
}
