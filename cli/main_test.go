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
	"github.com/Fitbit/smartling/test"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Setenv(envVar(projectFileFlagName), "testdata/smartling.yml")

	os.Exit(m.Run())
}

func TestRun(t *testing.T) {
	ts := test.MockServer()

	defer ts.Close()

	a := assert.New(t)

	os.Args = []string{"cli", "pull"}

	_, err := test.CaptureStdout(func() error {
		main()

		return nil
	})

	a.NoError(err)
}
