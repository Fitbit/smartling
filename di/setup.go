// Copyright 2017, Fitbit, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and limitations under the License.
package di

import (
	"github.com/Fitbit/smartling/repository"
	"github.com/Fitbit/smartling/rest"
	"github.com/Fitbit/smartling/service"
	"github.com/facebookgo/inject"
)

func Setup(opts *Options) (*Container, error) {
	var g inject.Graph

	c := Container{}

	err := g.Provide(
		&inject.Object{
			Value: &c,
		},
		&inject.Object{
			Value: &repository.YmlProjectConfigRepository{
				Filename: opts.Filename,
			},
			Name: "YmlProjectConfigRepository",
		},
		&inject.Object{
			Value: &service.DefaultProjectConfigService{},
			Name:  "DefaultProjectConfigService",
		},
		&inject.Object{
			Value: &service.DefaultAuthService{},
			Name:  "DefaultAuthService",
		},
		&inject.Object{
			Value: &service.DefaultFileService{},
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

	return &c, err
}
