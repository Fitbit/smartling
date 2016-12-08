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
	"github.com/mdreizin/smartling/model"
	"github.com/mdreizin/smartling/service"
	"gopkg.in/urfave/cli.v1"
)

func persistAuthTokenAction(c *cli.Context) (err error) {
	if c.App.Metadata[containerKey] != nil {
		container := c.App.Metadata[containerKey].(*service.Container)

		if c.App.Metadata[authTokenKey] != nil {
			authToken := c.App.Metadata[authTokenKey].(*model.AuthToken)
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
