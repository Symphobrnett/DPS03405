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

package legacy

import (
	"flag"

	"github.com/edgexfoundry/edgex-go/internal/security/secrets/option/contract"
	"github.com/edgexfoundry/edgex-go/internal/security/secrets/option/helper"
)

const CommandLegacy = "legacy"

type Command struct {
	configFile string
	helper     *helper.Helper
}

func NewCommand(flags *FlagSet, helper *helper.Helper) (*Command, *flag.FlagSet) {
	return &Command{
			configFile: flags.configFile,
			helper:     helper,
		},
		flags.flagSet
}

func (c *Command) Execute() (statusCode int, err error) {
	err = c.helper.GenTLSAssets(c.configFile)
	if err != nil {
		statusCode = contract.StatusCodeExitWithError
	} else {
		statusCode = contract.StatusCodeExitNormal
	}
	return
}
