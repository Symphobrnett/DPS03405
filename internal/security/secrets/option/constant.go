//
// Copyright (c) 2019 Intel Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
// in compliance with the License. You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed under the License
// is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
// or implied. See the License for the specific language governing permissions and limitations under
// the License.
//
// SPDX-License-Identifier: Apache-2.0'
//

package option

import (
	"errors"
)

const (
	// SecuritySecretsSetup is the name of this module
	SecuritySecretsSetup = "security-secrets-setup"

	PkiSetupVaultJSON             = "pkisetup-vault.json"
	PkiSetupKongJSON              = "pkisetup-kong.json"
	PkiInitScratchDir             = "scratch"
	PkiInitGeneratedDir           = "generated"
	ResourceDirName               = "res"
	ConfigTomlFile                = "configuration.toml"
	EnvXdgRuntimeDir              = "XDG_RUNTIME_DIR"
	DefaultWorkDir                = "/tmp"
	PkiInitBaseDir                = "edgex/security-secrets-setup"
	DefaultPkiCacheDir            = "/etc/edgex/pki"
	DefaultPkiDeployDir           = "/run/edgex/secrets"
	TlsSecretFileName             = "server.key"
	TlsCertFileName               = "server.crt"
	CaCertFileName                = "ca.pem"
	PkiInitFilePerServiceComplete = ".security-secrets-secrets.complete"

	// service name section:
	CaServiceName    = "ca"
	VaultServiceName = "edgex-vault"
	KongServiceName  = "edgex-kong"

	// ExitNormal exit code
	ExitNormal int = 0
	// ExitWithError is exit code for error
	ExitWithError int = 2
)

var (
	ErrCacheNotChangeAfter = errors.New("PKI cache cannot be changed after it was cached previously")
)
