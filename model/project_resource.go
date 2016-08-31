package model

import (
	"bytes"
	"github.com/Masterminds/sprig"
	"github.com/mattn/go-zglob"
	"html/template"
	"math"
	"path"
	"regexp"
	"strings"
)

type ProjectResource struct {
	Type             string       `yaml:"Type"`
	PathGlob         string       `yaml:"PathGlob"`
	PathExpression   string       `yaml:"PathExpression"`
	PathExclude      []string     `yaml:"PathExclude,omitempty"`
	AuthorizeContent bool         `yaml:"AuthorizeContent"`
	Directives       DirectiveMap `yaml:"Directives,omitempty"`
}

func (r *ProjectResource) PathFor(filename string, locale string) (string, error) {
	funcMap := sprig.FuncMap()

	t, err := template.New("PathForTemplate").Funcs(funcMap).Parse(r.PathExpression)

	if err != nil {
		return "", err
	}

	dir := path.Dir(filename)
	base := path.Base(filename)
	ext := path.Ext(filename)

	data := struct {
		Path   string
		Dir    string
		Base   string
		Name   string
		Ext    string
		Locale string
	}{
		filename,
		dir,
		base,
		strings.TrimSuffix(base, ext),
		ext,
		locale,
	}

	wr := bytes.NewBufferString("")

	err = t.Execute(wr, data)

	return wr.String(), err
}

func (r *ProjectResource) Files() []string {
	allFiles, _ := zglob.Glob(r.PathGlob)

	if len(r.PathExclude) > 0 {
		files := []string{}

		for _, name := range allFiles {
			include := false

			for _, pattern := range r.PathExclude {
				matched, err := regexp.MatchString(pattern, name)

				if matched && err == nil {
					include = false

					break
				} else {
					include = true
				}
			}

			if include {
				files = append(files, name)
			}
		}

		return files
	}

	return allFiles
}

func (r *ProjectResource) LimitFiles(limit int) [][]string {
	files := r.Files()
	pages := [][]string{}

	if limit < 0 || len(files) < limit {
		pages = append(pages, files)
	} else if limit > 0 {
		p := int(math.Ceil(float64(float32(len(files)) / float32(limit))))

		for i := 0; i < p; i++ {
			var page []string

			low := limit * i
			high := limit * (i + 1)

			if i == 0 {
				page = files[0:limit]
			} else if i < p-1 {
				page = files[low:high]
			} else {
				page = files[low:]
			}

			pages = append(pages, page)
		}
	}

	return pages
}
