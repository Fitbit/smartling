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
	"github.com/Fitbit/smartling/di"
	"github.com/Fitbit/smartling/model"
	"gopkg.in/urfave/cli.v1"
)

func persistAuthTokenAction(c *cli.Context) (err error) {
	if c.GlobalBool(saveAccessTokenFlagName) && c.App.Metadata[containerMetadataKey] != nil {
		container := c.App.Metadata[containerMetadataKey].(*di.Container)

		if c.App.Metadata[authTokenMetadataKey] != nil {
			authToken := c.App.Metadata[authTokenMetadataKey].(*model.AuthToken)
			projectConfig := model.ProjectConfig{
				AuthToken: model.AuthToken{
					AccessToken: authToken.RefreshToken,
				},
			}

			err = container.ProjectConfigService.UpdateConfig(&projectConfig)
		}
	}

	return err
}
