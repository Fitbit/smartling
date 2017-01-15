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
	"github.com/Fitbit/smartling/service"
	"gopkg.in/go-playground/pool.v3"
)

func pullJob(req *pullRequest) pool.WorkFunc {
	return func(wu pool.WorkUnit) (interface{}, error) {
		if wu.IsCancelled() {
			return nil, nil
		}

		uris := []string{}

		for _, path := range req.Files {
			uris = append(uris, req.Config.FileURI(path))
		}

		params := &service.FilePullParams{
			ProjectID:              req.Config.Project.ID,
			FileURIs:               uris,
			LocaleIDs:              req.Locales,
			RetrievalType:          req.RetrievalType,
			IncludeOriginalStrings: req.IncludeOriginalStrings,
			AuthToken:              req.AuthToken,
		}

		files, err := req.FileService.Pull(params)

		return &pullResponse{
			Files:   files,
			Request: req,
		}, err
	}
}
