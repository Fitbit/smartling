package model

import (
	"bytes"
	"github.com/Masterminds/sprig"
	"html/template"
	"path"
	"path/filepath"
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

	data := struct {
		Path   string
		Dir    string
		Base   string
		Ext    string
		Locale string
	}{
		filename,
		path.Dir(filename),
		path.Base(filename),
		path.Ext(filename),
		locale,
	}

	wr := bytes.NewBufferString("")

	err = t.Execute(wr, data)

	return wr.String(), err
}

func (r *ProjectResource) Files() []string {
	allFiles, _ := filepath.Glob(r.PathGlob)

	if len(r.PathExclude) > 0 {
		files := []string{}

		for _, name := range allFiles {
			include := false

			for _, pattern := range r.PathExclude {
				matched, err := filepath.Match(pattern, name)

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
