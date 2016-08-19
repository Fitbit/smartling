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
