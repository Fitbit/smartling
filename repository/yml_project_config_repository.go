// Copyright 2016, Fitbit, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and limitations under the License.
package repository

import (
	"github.com/mdreizin/smartling/model"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type YmlProjectConfigRepository struct {
	Filename string
}

func (r *YmlProjectConfigRepository) GetConfig() (*model.ProjectConfig, error) {
	var (
		err   error
		bytes []byte
	)

	config := &model.ProjectConfig{}

	if bytes, err = ioutil.ReadFile(r.Filename); err == nil {
		err = yaml.Unmarshal(bytes, &config)
	}

	return config, err
}

func (r *YmlProjectConfigRepository) UpdateConfig(delta *model.ProjectConfig) error {
	var (
		src   *model.ProjectConfig
		err   error
		bytes []byte
	)

	if src, err = r.GetConfig(); err == nil {
		if err = src.Merge(delta); err == nil {
			if bytes, err = yaml.Marshal(src); err == nil {
				err = ioutil.WriteFile(r.Filename, bytes, 0644)
			}
		}
	}

	return err
}
