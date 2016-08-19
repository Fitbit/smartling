package service

import (
	"github.com/mdreizin/smartling/model"
	"github.com/mdreizin/smartling/repository"
)

type DefaultProjectConfigService struct {
	ProjectConfigRepository repository.ProjectConfigRepository `inject:"YmlProjectConfigRepository"`
}

func (s *DefaultProjectConfigService) GetConfig() (*model.ProjectConfig, error) {
	return s.ProjectConfigRepository.GetConfig()
}

func (s *DefaultProjectConfigService) UpdateConfig(delta *model.ProjectConfig) error {
	return s.ProjectConfigRepository.UpdateConfig(delta)
}
