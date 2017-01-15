// Copyright 2016, Fitbit, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and limitations under the License.
package main

import (
	"fmt"
	"github.com/Fitbit/smartling/model"
	"gopkg.in/urfave/cli.v1"
)

func validateProjectConfigAction(c *cli.Context) error {
	var err error

	invalidOptions := []string{}

	projectConfig := c.App.Metadata[projectConfigMetadataKey].(*model.ProjectConfig)

	if projectConfig.Project.ID == "" {
		invalidOptions = append(invalidOptions, projectIDFlagName)
	}

	if projectConfig.UserToken.ID == "" {
		invalidOptions = append(invalidOptions, userTokenIDFlagName)
	}

	if projectConfig.UserToken.Secret == "" {
		invalidOptions = append(invalidOptions, userTokenSecretFlagName)
	}

	if len(invalidOptions) > 0 {
		err = fmt.Errorf(`The following "GLOBAL OPTIONS" are required %v`, invalidOptions)
	}

	return err
}
