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
	"archive/zip"
	"bytes"
	"github.com/google/go-querystring/query"
	"github.com/mdreizin/smartling/model"
	"github.com/mdreizin/smartling/rest"
	"gopkg.in/go-playground/pool.v3"
	"gopkg.in/resty.v0"
	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

type DefaultFileService struct {
	Client *resty.Client `inject:"DefaultRestClient"`
}

func (s *DefaultFileService) Pull(params *FilePullParams) ([]*model.File, error) {
	var (
		q      url.Values
		err    error
		resp   *resty.Response
		_url   string
		reader *zip.Reader
	)

	files := []*model.File{}

	p := pool.New()

	defer p.Close()

	batch := p.Batch()

	if _url, err = rest.GenerateURL(rest.FilePullURL, &params); err == nil {
		if q, err = query.Values(params); err == nil {
			q.Add("fileNameMode", "UNCHANGED")
			q.Add("localeMode", "LOCALE_IN_NAME")

			req := s.Client.R().SetMultiValueQueryParams(q).SetAuthToken(params.AuthToken)

			if resp, err = req.Get(_url); err == nil {
				body := resp.Body()

				go func(params *FilePullParams) {
					if reader, err = zip.NewReader(bytes.NewReader(body), int64(len(body))); err == nil {
						for _, file := range reader.File {
							batch.Queue(s.extractFileJob(file, params.LocaleIDs))
						}
					}

					batch.QueueComplete()
				}(params)
			}
		}
	}

	for result := range batch.Results() {
		resp := result.Value().(*model.File)

		if err := result.Error(); err == nil {
			files = append(files, resp)
		}
	}

	return files, err
}

func (s *DefaultFileService) Push(params *FilePushParams) (*model.FileStats, error) {
	var (
		q        url.Values
		err      error
		resp     *resty.Response
		_url     string
		filename string
		reader   *os.File
	)

	stats := &model.FileStats{}

	if _url, err = rest.GenerateURL(rest.FilePushURL, &params); err == nil {
		if q, err = query.Values(params); err == nil {
			if filename, err = filepath.Abs(params.FilePath); err == nil {
				if reader, err = os.Open(filename); err == nil {
					req := s.Client.R().
						SetResult(rest.Model{}).
						SetError(rest.Model{}).
						SetAuthToken(params.AuthToken).
						SetFileReader("file", "", reader)

					for p, v := range q {
						for _, pv := range v {
							req.FormData.Add(p, pv)
						}
					}

					req.SetFormData(params.Directives)

					if resp, err = req.Post(_url); err == nil {
						err = rest.Result(resp, &stats)
					}
				}
			}
		}
	}

	return stats, err
}

func (s *DefaultFileService) extractFile(zf *zip.File, locales []string) (*model.File, error) {
	var err error

	file := &model.File{}

	if rc, err := zf.Open(); err == nil {
		if b, err := ioutil.ReadAll(rc); err == nil {
			for _, locale := range locales {
				name := zf.Name
				localeSuffix := "_" + locale

				if strings.Contains(name, localeSuffix) {
					file.Path = strings.Replace(name, localeSuffix, "", -1)
					file.LocaleID = locale
					break
				}
			}

			file.Content = b
		}

		rc.Close()
	}

	return file, err
}

func (s *DefaultFileService) extractFileJob(zf *zip.File, locales []string) pool.WorkFunc {
	return func(wu pool.WorkUnit) (interface{}, error) {
		if wu.IsCancelled() {
			return nil, nil
		}

		return s.extractFile(zf, locales)
	}
}
