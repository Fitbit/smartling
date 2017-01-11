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
	"github.com/Fitbit/smartling/service"
	"github.com/fatih/color"
	"gopkg.in/go-playground/pool.v3"
	"gopkg.in/urfave/cli.v1"
	"runtime"
	"time"
)

var pushCommand = cli.Command{
	Name:  "push",
	Usage: "Uploads translations",
	Before: func(c *cli.Context) error {
		return invokeActions([]action{
			ensureMetadataAction,
			injectContainerAction,
			injectProjectConfigAction,
			validateProjectConfigAction,
			injectAuthTokenAction,
		}, c)
	},
	Action: func(c *cli.Context) error {
		defer elapsedTime(time.Now())

		p := pool.NewLimited(uint(runtime.NumCPU()))

		defer p.Close()

		batch := p.Batch()

		container := c.App.Metadata[containerKey].(*service.Container)
		authToken := c.App.Metadata[authTokenKey].(*model.AuthToken)
		projectConfig := c.App.Metadata[projectConfigKey].(*model.ProjectConfig)

		go func() {
			for _, resource := range projectConfig.Resources {
				for _, path := range resource.Files() {
					batch.Queue(pushJob(&pushRequest{
						Path:        path,
						Resource:    &resource,
						Config:      projectConfig,
						AuthToken:   authToken.AccessToken,
						FileService: container.FileService,
					}))
				}
			}

			batch.QueueComplete()
		}()

		i := 0

		for result := range batch.Results() {
			resp := result.Value().(*pushResponse)

			if err := result.Error(); err != nil {
				logError(fmt.Sprintf("%s has error %s", resp.Params.FilePath, color.RedString(err.Error())))
			} else {
				logInfo(fmt.Sprintf("%s {Override=%t Strings=%d Words=%d}", color.MagentaString(resp.Params.FilePath), resp.Stats.OverWritten, resp.Stats.StringCount, resp.Stats.WordCount))
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
