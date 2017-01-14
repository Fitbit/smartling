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
	"github.com/fatih/color"
	"gopkg.in/urfave/cli.v1"
)

var listCommand = cli.Command{
	Name: "list",
	Aliases: []string{
		"ls",
	},
	Usage: "Shows a list of local translations",
	Before: func(c *cli.Context) error {
		return invokeActions([]action{
			injectDiContainerAction,
			injectProjectConfigAction,
		}, c)
	},
	Action: func(c *cli.Context) error {
		projectConfig := c.App.Metadata[projectConfigMetadataKey].(*model.ProjectConfig)
		i := 0

		for _, resource := range projectConfig.Resources {
			logInfo(fmt.Sprintf("Using {PathGlob=%v PathExclude=%v}", resource.PathGlob, resource.PathExclude))

			for _, v := range resource.Files() {
				logInfo(color.MagentaString(v))
				i++
			}
		}

		logInfo(fmt.Sprintf("%d files", i))

		return nil
	},
	After: func(c *cli.Context) error {
		return invokeActions([]action{
			persistAuthTokenAction,
		}, c)
	},
}
