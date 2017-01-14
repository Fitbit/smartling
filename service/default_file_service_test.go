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
	"github.com/Fitbit/smartling/rest"
	"github.com/Fitbit/smartling/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDefaultFileService_Push(t *testing.T) {
	ts := test.MockServer()

	defer ts.Close()

	fileService := DefaultFileService{
		Client: rest.Client(false),
	}

	stats, err := fileService.Push(&FilePushParams{
		ProjectID:  "projectId",
		FilePath:   "../test/testdata/foo/en-US.json",
		Directives: map[string]string{},
		AuthToken:  "authToken",
	})

	a := assert.New(t)

	a.NoError(err)
	a.EqualValues(&model.FileStats{
		OverWritten: true,
		StringCount: 10,
		WordCount:   10,
	}, stats)
}

func TestDefaultFileService_Pull(t *testing.T) {
	ts := test.MockServer()

	defer ts.Close()

	fileService := DefaultFileService{
		Client: rest.Client(false),
	}

	files, err := fileService.Pull(&FilePullParams{
		ProjectID: "projectId",
		FileURIs:  []string{},
		LocaleIDs: []string{"de-DE"},
	})

	a := assert.New(t)

	a.NoError(err)
	a.NotNil(files)
	a.Len(files, 1)
}
