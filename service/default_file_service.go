package service

import (
	"archive/zip"
	"bytes"
	"github.com/google/go-querystring/query"
	"github.com/mdreizin/smartling/model"
	"github.com/mdreizin/smartling/rest"
	"gopkg.in/resty.v0"
	"io/ioutil"
	"net/url"
	"path/filepath"
	"strings"
)

func extractFile(zf *zip.File, locales []string) (file *model.File, err error) {
	file = &model.File{}

	if rc, err := zf.Open(); err == nil {
		if b, err := ioutil.ReadAll(rc); err == nil {
			for _, locale := range locales {
				name := zf.Name
				localeSuffix := "_" + locale

				if strings.Contains(name, localeSuffix) {
					file.Path = strings.Replace(name, localeSuffix, "", -1)
					file.LocaleID = locale
					break
				}
			}

			file.Content = b
		}

		rc.Close()
	}

	return file, err
}

type DefaultFileService struct{}

func (s *DefaultFileService) Pull(params *FilePullParams) ([]*model.File, error) {
	var (
		q      url.Values
		err    error
		resp   *resty.Response
		_url   string
		reader *zip.Reader
	)

	files := []*model.File{}

	if _url, err = rest.DynamicURL(rest.FilePullURL, &params); err == nil {
		if q, err = query.Values(params); err == nil {
			q.Add("fileNameMode", "UNCHANGED")
			q.Add("localeMode", "LOCALE_IN_NAME")

			req := resty.R().SetMultiValueQueryParams(q).SetAuthToken(params.AuthToken)

			if resp, err = req.Get(_url); err == nil {
				body := resp.Body()

				if reader, err = zip.NewReader(bytes.NewReader(body), int64(len(body))); err == nil {
					for _, zf := range reader.File {
						if file, err := extractFile(zf, params.LocaleIDs); err == nil {
							files = append(files, file)
						}
					}
				}
			}
		}
	}

	return files, err
}

func (s *DefaultFileService) Push(params *FilePushParams) (*model.FileStats, error) {
	var (
		q        url.Values
		err      error
		resp     *resty.Response
		_url     string
		filename string
		content  []byte
	)

	stats := &model.FileStats{}

	if _url, err = rest.DynamicURL(rest.FilePushURL, &params); err == nil {
		if q, err = query.Values(params); err == nil {
			if filename, err = filepath.Abs(params.FilePath); err == nil {
				if content, err = ioutil.ReadFile(filename); err == nil {
					req := resty.R().
						SetResult(rest.Model{}).
						SetError(rest.Model{}).
						SetAuthToken(params.AuthToken).
						SetFileReader("file", "", bytes.NewReader(content))

					for p, v := range q {
						for _, pv := range v {
							req.FormData.Add(p, pv)
						}
					}

					req.SetFormData(params.Directives)

					if resp, err = req.Post(_url); err == nil {
						err = rest.Result(resp, &stats)
					}
				}
			}
		}
	}

	return stats, err
}
