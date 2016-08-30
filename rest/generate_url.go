package rest

import (
	"bytes"
	"text/template"
)

func GenerateURL(url string, data interface{}) (string, error) {
	t, err := template.New("GenerateURLTemplate").Parse(url)
	wr := bytes.NewBufferString("")

	if err == nil {
		err = t.Execute(wr, data)
	}

	return wr.String(), err
}
