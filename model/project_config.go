package model

import (
	"fmt"
	"github.com/imdario/mergo"
	"io/ioutil"
	"path"
	"path/filepath"
	"strings"
)

type ProjectConfig struct {
	UserToken `yaml:",inline"`
	Project   `yaml:",inline"`
	AuthToken `yaml:",inline"`
	Resources []ProjectResource `yaml:"Files,omitempty"`
	Locales   map[string]string `yaml:"Locales,omitempty"`
}

func (c *ProjectConfig) Merge(delta *ProjectConfig) error {
	return mergo.MapWithOverwrite(c, delta)
}

func (p *ProjectConfig) LocaleFor(localeID string) string {
	locale := p.Locales[localeID]

	if locale != "" {
		return p.Locales[localeID]
	}

	return localeID
}

func (p *ProjectConfig) FileURI(filename string) string {
	if p.Alias != "" {
		return path.Join(p.Alias, filename)
	}

	return filename
}

func (p *ProjectConfig) FilePath(filename string) string {
	return strings.Trim(strings.TrimPrefix(filename, p.Alias), fmt.Sprintf("%c", filepath.Separator))
}

func (p *ProjectConfig) SaveFile(file *File, resource *ProjectResource) error {
	var (
		err      error
		filename string
	)

	locale := p.LocaleFor(file.LocaleID)

	if filename, err = resource.PathFor(p.FilePath(file.Path), locale); err == nil {
		if filename, err = filepath.Abs(filename); err == nil {
			err = ioutil.WriteFile(filename, file.Content, 0644)
		}
	}

	return err
}
