package test

import (
	"archive/zip"
	"bytes"
	"fmt"
	"github.com/mdreizin/smartling/rest"
	"net/http"
	"net/http/httptest"
	"strings"
)

const (
	contentTypeHeader = "Content-Type"
	jsonContentType   = "application/json"
)

func resp(code string, data string) string {
	return fmt.Sprintf(`{"response":{"code":"%s","data":%s}}`, code, data)
}

func ok(data string) string {
	return resp("SUCCESS", data)
}

func zipFile() *bytes.Buffer {
	buf := new(bytes.Buffer)
	w := zip.NewWriter(buf)

	var files = []struct {
		Name, Body string
	}{
		{"../test/testdata/foo/en-US_ru-RU.json", "{}"},
	}

	for _, file := range files {
		f, _ := w.Create(file.Name)

		f.Write([]byte(file.Body))
	}

	w.Close()

	return buf
}

var (
	filePushURL = strings.Replace(rest.FilePushURL, "{{ .ProjectID }}", "projectId", -1)
	filePullURL = strings.Replace(rest.FilePullURL, "{{ .ProjectID }}", "projectId", -1)
)

func MockServer() *httptest.Server {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, rest.AuthenticateURL) || strings.HasSuffix(r.URL.Path, rest.AuthenticateRefreshURL) {
			w.Header().Set(contentTypeHeader, jsonContentType)
			fmt.Fprint(w, ok(`{"accessToken":"accessToken","refreshToken":"refreshToken"}`))
		} else if strings.HasSuffix(r.URL.Path, filePushURL) {
			w.Header().Set(contentTypeHeader, jsonContentType)
			fmt.Fprint(w, ok(`{"overWritten":true,"stringCount":10,"wordCount":10}`))
		} else if strings.HasSuffix(r.URL.Path, filePullURL) {
			zf := zipFile()

			w.Write(zf.Bytes())
		}
	}))

	rest.BaseURL = ts.URL

	return ts
}
