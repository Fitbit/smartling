package model

import (
	"fmt"
	"github.com/imdario/mergo"
	"gopkg.in/go-playground/pool.v3"
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

func (c *ProjectConfig) LocaleFor(localeID string) string {
	locale := c.Locales[localeID]

	if locale != "" {
		return c.Locales[localeID]
	}

	return localeID
}

func (c *ProjectConfig) FileURI(filename string) string {
	if c.Alias != "" {
		return path.Join(c.Alias, filename)
	}

	return filename
}

func (c *ProjectConfig) FilePath(filename string) string {
	return strings.Trim(strings.TrimPrefix(filename, c.Alias), fmt.Sprintf("%c", filepath.Separator))
}

func (c *ProjectConfig) SaveFile(file *File, resource *ProjectResource) error {
	var (
		err      error
		filename string
	)

	locale := c.LocaleFor(file.LocaleID)

	if filename, err = resource.PathFor(c.FilePath(file.Path), locale); err == nil {
		if filename, err = filepath.Abs(filename); err == nil {
			err = ioutil.WriteFile(filename, file.Content, 0644)
		}
	}

	return err
}

func (c *ProjectConfig) SaveAllFiles(files []*File, resource *ProjectResource) {
	p := pool.New()

	defer p.Close()

	batch := p.Batch()

	go func() {
		for _, file := range files {
			batch.Queue(c.saveFileJob(file, resource))
		}

		batch.QueueComplete()
	}()

	batch.WaitAll()
}

func (c *ProjectConfig) saveFileJob(file *File, resource *ProjectResource) pool.WorkFunc {
	return func(wu pool.WorkUnit) (interface{}, error) {
		if wu.IsCancelled() {
			return nil, nil
		}

		return nil, c.SaveFile(file, resource)
	}
}
