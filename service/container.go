// Copyright 2016, Fitbit, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and limitations under the License.
package service

import (
	"github.com/Fitbit/smartling/repository"
	"github.com/Fitbit/smartling/rest"
	"github.com/facebookgo/inject"
)

type Container struct {
	ProjectConfigService ProjectConfigService `inject:"DefaultProjectConfigService"`
	AuthService          AuthService          `inject:"DefaultAuthService"`
	FileService          FileService          `inject:"DefaultFileService"`
}

func (c *Container) SetUp(filename string) error {
	var g inject.Graph

	err := g.Provide(
		&inject.Object{
			Value: c,
		},
		&inject.Object{
			Value: &repository.YmlProjectConfigRepository{
				Filename: filename,
			},
			Name: "YmlProjectConfigRepository",
		},
		&inject.Object{
			Value: &DefaultProjectConfigService{},
			Name:  "DefaultProjectConfigService",
		},
		&inject.Object{
			Value: &DefaultAuthService{},
			Name:  "DefaultAuthService",
		},
		&inject.Object{
			Value: &DefaultFileService{},
			Name:  "DefaultFileService",
		},
		&inject.Object{
			Value: rest.Client(),
			Name:  "DefaultRestClient",
		},
	)

	if err == nil {
		err = g.Populate()
	}

	return err
}
