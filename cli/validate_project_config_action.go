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
	"errors"
	"github.com/mdreizin/smartling/model"
	"gopkg.in/urfave/cli.v1"
	"strings"
)

func validateProjectConfigAction(c *cli.Context) error {
	var err error

	issues := []string{}

	projectConfig := c.App.Metadata[projectConfigKey].(*model.ProjectConfig)

	if projectConfig.Project.ID == "" {
		issues = append(issues, "project-id is required")
	}

	if projectConfig.UserToken.ID == "" {
		issues = append(issues, "user-id is required")
	}

	if projectConfig.UserToken.Secret == "" {
		issues = append(issues, "user-secret is required")
	}

	if len(issues) > 0 {
		err = errors.New(strings.Join(issues, "\n"))
	}

	return err
}
