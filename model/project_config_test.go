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
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProjectConfig_Merge(t *testing.T) {
	a := assert.New(t)
	src := ProjectConfig{
		AuthToken: AuthToken{
			AccessToken: "accessToken1",
		},
		UserToken: UserToken{
			ID: "id1",
		},
		Project: Project{
			Alias: "projectAlias1",
		},
	}
	dst := ProjectConfig{
		AuthToken: AuthToken{
			AccessToken: "accessToken2",
		},
		UserToken: UserToken{
			ID: "id2",
		},
		Project: Project{
			Alias: "projectAlias2",
		},
	}

	err := src.Merge(&dst)

	a.NoError(err)
	a.EqualValues(dst, src)
}

func TestProjectConfig_Locale(t *testing.T) {
	a := assert.New(t)
	p := ProjectConfig{
		Locales: map[string]string{
			"en-US": "en_US",
			"ru-RU": "ru",
		},
	}

	a.EqualValues("en_US", p.Locale("en-US"))
	a.EqualValues("ru", p.Locale("ru-RU"))
	a.EqualValues("de-DE", p.Locale("de-DE"))
}

func TestProjectConfig_FileURI(t *testing.T) {
	a := assert.New(t)
	p1 := ProjectConfig{
		Project: Project{
			Alias: "testdata",
		},
	}
	p2 := ProjectConfig{}

	a.EqualValues("testdata/foo.json", p1.FileURI("foo.json"))
	a.EqualValues("foo.json", p2.FileURI("foo.json"))
}

func TestProjectConfig_FilePath(t *testing.T) {
	a := assert.New(t)
	p1 := ProjectConfig{
		Project: Project{
			Alias: "testdata",
		},
	}
	p2 := ProjectConfig{}

	a.EqualValues("foo.json", p1.FilePath("testdata/foo.json"))
	a.EqualValues("foo.json", p2.FilePath("foo.json"))
	a.EqualValues("testdata/foo.json", p2.FilePath("testdata/foo.json"))
}
