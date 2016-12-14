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
	"github.com/Fitbit/smartling/model"
	"github.com/Fitbit/smartling/repository"
	"github.com/Fitbit/smartling/test"
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"testing"
)

func TestDefaultProjectConfigService_GetConfig(t *testing.T) {
	a := assert.New(t)
	src, _ := filepath.Abs("../repository/testdata/smartling.yml")
	projectConfigService := DefaultProjectConfigService{
		ProjectConfigRepository: &repository.YmlProjectConfigRepository{
			Filename: src,
		},
	}
	conf, err := projectConfigService.GetConfig()

	a.NoError(err)
	a.NotNil(conf)
}

func TestDefaultProjectConfigService_UpdateConfig(t *testing.T) {
	a := assert.New(t)
	src, _ := filepath.Abs("../repository/testdata/smartling.yml")
	dst, _ := filepath.Abs("../repository/testdata/.smartling.yml")

	projectConfigService := DefaultProjectConfigService{
		ProjectConfigRepository: &repository.YmlProjectConfigRepository{
			Filename: dst,
		},
	}

	test.CopyFile(src, dst, func() {
		d := model.ProjectConfig{
			AuthToken: model.AuthToken{
				AccessToken: "accessToken",
			},
		}

		err := projectConfigService.UpdateConfig(&d)

		a.NoError(err)
	})
}
