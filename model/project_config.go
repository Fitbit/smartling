// Copyright 2016, Fitbit, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and limitations under the License.
package model

import (
	"fmt"
	"github.com/imdario/mergo"
	"gopkg.in/go-playground/pool.v3"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type ProjectConfig struct {
	UserToken `yaml:",inline"`
	Project   `yaml:",inline"`
	AuthToken `yaml:",inline"`
	Resources []ProjectResource `yaml:"Files,omitempty"`
	Locales   map[string]string `yaml:"Locales,omitempty"`
}

func (c *ProjectConfig) Merge(delta *ProjectConfig) error {
	return mergo.MapWithOverwrite(c, delta)
}

func (c *ProjectConfig) Locale(localeID string) string {
	locale := c.Locales[localeID]

	if locale != "" {
		return c.Locales[localeID]
	}

	return localeID
}

func (c *ProjectConfig) FileURI(filename string) string {
	if c.Alias != "" {
		return path.Join(c.Alias, filename)
	}

	return filename
}

func (c *ProjectConfig) FilePath(filename string) string {
	return strings.Trim(strings.TrimPrefix(filename, c.Alias), fmt.Sprintf("%c", filepath.Separator))
}

func (c *ProjectConfig) SaveFile(file *File, resource *ProjectResource) error {
	var (
		err      error
		filename string
	)

	locale := c.Locale(file.LocaleID)

	if filename, err = resource.FilePath(c.FilePath(file.Path), locale); err == nil {
		if filename, err = filepath.Abs(filename); err == nil {
			dir := filepath.Dir(filename)

			if err = os.MkdirAll(dir, 0777); err == nil {
				err = ioutil.WriteFile(filename, file.Content, 0644)
			}
		}
	}

	return err
}

func (c *ProjectConfig) SaveAllFiles(files []*File, resource *ProjectResource) map[*File]error {
	errs := map[*File]error{}
	p := pool.New()

	defer p.Close()

	batch := p.Batch()

	go func() {
		for _, file := range files {
			batch.Queue(c.saveFileWorker(file, resource))
		}

		batch.QueueComplete()
	}()

	for results := range batch.Results() {
		result := results.Value().(*File)

		if err := results.Error(); err != nil && result != nil {
			errs[result] = err
		}
	}

	return errs
}

func (c *ProjectConfig) saveFileWorker(file *File, resource *ProjectResource) pool.WorkFunc {
	return func(wu pool.WorkUnit) (interface{}, error) {
		if wu.IsCancelled() {
			return nil, nil
		}

		return file, c.SaveFile(file, resource)
	}
}
