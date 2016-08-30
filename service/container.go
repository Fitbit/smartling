package service

import (
	"github.com/facebookgo/inject"
	"github.com/mdreizin/smartling/repository"
	"github.com/mdreizin/smartling/rest"
)

type Container struct {
	ProjectConfigService ProjectConfigService `inject:"DefaultProjectConfigService"`
	AuthService          AuthService          `inject:"DefaultAuthService"`
	FileService          FileService          `inject:"DefaultFileService"`
}

func (c *Container) SetUp(filename string) error {
	var g inject.Graph

	err := g.Provide(
		&inject.Object{
			Value: c,
		},
		&inject.Object{
			Value: &repository.YmlProjectConfigRepository{
				Filename: filename,
			},
			Name: "YmlProjectConfigRepository",
		},
		&inject.Object{
			Value: &DefaultProjectConfigService{},
			Name:  "DefaultProjectConfigService",
		},
		&inject.Object{
			Value: &DefaultAuthService{},
			Name:  "DefaultAuthService",
		},
		&inject.Object{
			Value: &DefaultFileService{},
			Name:  "DefaultFileService",
		},
		&inject.Object{
			Value: rest.Client(),
			Name:  "DefaultRestClient",
		},
	)

	if err == nil {
		err = g.Populate()
	}

	return err
}
