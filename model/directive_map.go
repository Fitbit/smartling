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
	"strings"
)

const DirectivePrefix = "smartling."

type DirectiveMap map[string]string

func (d DirectiveMap) WithPrefix() map[string]string {
	m := map[string]string{}

	for key, value := range d {
		if !strings.HasPrefix(key, DirectivePrefix) {
			m[DirectivePrefix+key] = value
		} else {
			m[key] = value
		}
	}

	return m
}
