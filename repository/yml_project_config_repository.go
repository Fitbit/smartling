package repository

import (
	"github.com/mdreizin/smartling/model"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type YmlProjectConfigRepository struct {
	Filename string
}

func (r *YmlProjectConfigRepository) GetConfig() (*model.ProjectConfig, error) {
	var (
		err   error
		bytes []byte
	)

	config := &model.ProjectConfig{}

	if bytes, err = ioutil.ReadFile(r.Filename); err == nil {
		err = yaml.Unmarshal(bytes, &config)
	}

	return config, err
}

func (r *YmlProjectConfigRepository) UpdateConfig(delta *model.ProjectConfig) error {
	var (
		src   *model.ProjectConfig
		err   error
		bytes []byte
	)

	if src, err = r.GetConfig(); err == nil {
		if err = src.Merge(delta); err == nil {
			if bytes, err = yaml.Marshal(src); err == nil {
				err = ioutil.WriteFile(r.Filename, bytes, 0644)
			}
		}
	}

	return err
}
