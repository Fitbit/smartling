package rest

import (
	"bytes"
	"text/template"
)

func DynamicURL(url string, data interface{}) (string, error) {
	t, err := template.New("DynamicURLTemplate").Parse(StaticURL(url))
	wr := bytes.NewBufferString("")

	if err == nil {
		err = t.Execute(wr, data)
	}

	return wr.String(), err
}
