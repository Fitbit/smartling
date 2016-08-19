package repository

import "github.com/mdreizin/smartling/model"

type ProjectConfigRepository interface {
	GetConfig() (*model.ProjectConfig, error)

	UpdateConfig(*model.ProjectConfig) error
}
