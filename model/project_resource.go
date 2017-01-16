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
	"bytes"
	"github.com/Masterminds/sprig"
	"github.com/mattn/go-zglob"
	"html/template"
	"math"
	"path"
	"regexp"
	"strings"
)

type ProjectResource struct {
	Type             string       `yaml:"Type"`
	PathGlob         string       `yaml:"PathGlob"`
	PathExpression   string       `yaml:"PathExpression"`
	PathExclude      []string     `yaml:"PathExclude,omitempty"`
	AuthorizeContent bool         `yaml:"AuthorizeContent"`
	Directives       DirectiveMap `yaml:"Directives,omitempty"`
}

func (r *ProjectResource) FilePath(filename string, locale string) (string, error) {
	funcMap := sprig.FuncMap()

	t, err := template.New("FilePathTemplate").Funcs(funcMap).Parse(r.PathExpression)

	if err != nil {
		return "", err
	}

	dir := path.Dir(filename)
	base := path.Base(filename)
	ext := path.Ext(filename)

	data := struct {
		Path   string
		Dir    string
		Base   string
		Name   string
		Ext    string
		Locale string
	}{
		filename,
		dir,
		base,
		strings.TrimSuffix(base, ext),
		ext,
		locale,
	}

	wr := bytes.NewBufferString("")

	err = t.Execute(wr, data)

	return wr.String(), err
}

func (r *ProjectResource) Files() []string {
	allFiles, _ := zglob.Glob(r.PathGlob)

	if len(r.PathExclude) > 0 {
		files := []string{}

		for _, name := range allFiles {
			include := false

			for _, pattern := range r.PathExclude {
				matched, err := regexp.MatchString(pattern, name)

				if matched && err == nil {
					include = false

					break
				} else {
					include = true
				}
			}

			if include {
				files = append(files, name)
			}
		}

		return files
	}

	return allFiles
}

func (r *ProjectResource) BatchFiles(size int) [][]string {
	allFiles := r.Files()
	batch := [][]string{}

	if size < 0 || len(allFiles) < size {
		batch = append(batch, allFiles)
	} else if size > 0 {
		p := int(math.Ceil(float64(float32(len(allFiles)) / float32(size))))

		for i := 0; i < p; i++ {
			var files []string

			low := size * i
			high := size * (i + 1)

			if i == 0 {
				files = allFiles[0:size]
			} else if i < p-1 {
				files = allFiles[low:high]
			} else {
				files = allFiles[low:]
			}

			batch = append(batch, files)
		}
	}

	return batch
}
