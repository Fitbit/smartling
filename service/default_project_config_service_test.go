package service

import (
	"github.com/mdreizin/smartling/model"
	"github.com/mdreizin/smartling/repository"
	"github.com/mdreizin/smartling/test"
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"testing"
)

func TestDefaultProjectConfigService_GetConfig(t *testing.T) {
	a := assert.New(t)
	src, _ := filepath.Abs("../repository/testdata/smartling.yml")
	projectConfigService := DefaultProjectConfigService{
		ProjectConfigRepository: &repository.YmlProjectConfigRepository{
			Filename: src,
		},
	}
	conf, err := projectConfigService.GetConfig()

	a.NoError(err)
	a.NotNil(conf)
}

func TestDefaultProjectConfigService_UpdateConfig(t *testing.T) {
	a := assert.New(t)
	src, _ := filepath.Abs("../repository/testdata/smartling.yml")
	dst, _ := filepath.Abs("../repository/testdata/.smartling.yml")

	projectConfigService := DefaultProjectConfigService{
		ProjectConfigRepository: &repository.YmlProjectConfigRepository{
			Filename: dst,
		},
	}

	test.CopyFile(src, dst, func() {
		d := model.ProjectConfig{
			AuthToken: model.AuthToken{
				AccessToken: "accessToken",
			},
		}

		err := projectConfigService.UpdateConfig(&d)

		a.NoError(err)
	})
}
