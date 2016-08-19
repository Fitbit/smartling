package service

import "github.com/mdreizin/smartling/model"

type ProjectConfigService interface {
	GetConfig() (*model.ProjectConfig, error)

	UpdateConfig(*model.ProjectConfig) error
}
