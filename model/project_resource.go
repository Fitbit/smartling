package model

import (
	"bytes"
	"github.com/Masterminds/sprig"
	"github.com/mattn/go-zglob"
	"html/template"
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

func (r *ProjectResource) PartialFiles(limit int) [][]string {
	files := r.Files()
	partialFiles := [][]string{}

	if limit < 0 || len(files) < limit {
		partialFiles = append(partialFiles, files)
	} else if limit > 0 {
		q := []string{}
		i := 0

		for _, x := range files {
			i++
			q = append(q, x)

			if i == limit {
				partialFiles = append(partialFiles, q)

				q = []string{}
				i = 0
			}
		}
	}

	return partialFiles
}
