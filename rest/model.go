// Copyright 2016, Fitbit, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and limitations under the License.
package rest

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Model struct {
	Response struct {
		Code   string          `json:"code"`
		Data   json.RawMessage `json:"data,omitempty"`
		Errors []Error         `json:"errors,omitempty"`
	} `json:"response"`
}

func (m *Model) Data(data interface{}) (err error) {
	if m.Response.Data != nil {
		err = json.Unmarshal(m.Response.Data, &data)
	}

	return err
}

func (m *Model) Error() (err error) {
	errors := m.Response.Errors

	if len(errors) > 0 {
		messages := []string{}

		for _, e := range errors {
			messages = append(messages, e.Message)
		}

		err = fmt.Errorf(strings.Join(messages, "\n"))
	}

	return err
}

func (m *Model) IsOK() bool {
	return m.Response.Code == "SUCCESS"
}
