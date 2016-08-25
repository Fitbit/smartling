package repository

import (
	"github.com/mdreizin/smartling/model"
	"github.com/mdreizin/smartling/test"
	"github.com/stretchr/testify/assert"
	"path"
	"testing"
)

func TestYmlProjectConfigRepository_GetConfig(t *testing.T) {
	a := assert.New(t)
	projectConfigRepository := YmlProjectConfigRepository{
		Filename: "testdata/smartling.yml",
	}
	conf1, err := projectConfigRepository.GetConfig()
	conf2 := model.ProjectConfig{
		Project: model.Project{
			ID:    "projectId",
			Alias: "projectAlias",
		},
		UserToken: model.UserToken{
			ID:     "userId",
			Secret: "userSecret",
		},
		Resources: []model.ProjectResource{
			{
				Type:             "json",
				PathGlob:         "repository/testdata/*/en-US.json",
				PathExpression:   "{{ .Dir }}/{{ .Locale }}{{ .Ext }}",
				AuthorizeContent: true,
				Directives: model.DirectiveMap{
					"string_format": "NONE",
				},
			},
		},
		Locales: map[string]string{
			"ru-RU": "ru",
		},
	}

	a.NoError(err)
	a.EqualValues(&conf2, conf1)
}

func TestYmlProjectConfigRepository_GetConfig_ThrowsError(t *testing.T) {
	a := assert.New(t)
	projectConfigRepository := YmlProjectConfigRepository{Filename: "testdata/.smartling.yml"}
	_, err := projectConfigRepository.GetConfig()

	a.Error(err)
}

func TestYmlProjectConfigRepository_UpdateConfig(t *testing.T) {
	a := assert.New(t)
	src := path.Join("testdata", "smartling.yml")
	dst := path.Join("testdata", ".smartling.yml")
	projectConfigRepository := YmlProjectConfigRepository{
		Filename: dst,
	}

	test.CopyFile(src, dst, func() {
		c, _ := projectConfigRepository.GetConfig()
		d := model.ProjectConfig{
			AuthToken: model.AuthToken{
				AccessToken: "accessToken",
			},
		}

		projectConfigRepository.UpdateConfig(&d)

		l, _ := projectConfigRepository.GetConfig()

		a.EqualValues("", c.AccessToken)
		a.EqualValues("accessToken", l.AccessToken)
	})
}

func TestYmlProjectConfigRepository_UpdateConfig_ThrowsError(t *testing.T) {
	a := assert.New(t)
	projectConfigRepository := YmlProjectConfigRepository{Filename: "testdata/.smartling.yml"}
	err := projectConfigRepository.UpdateConfig(&model.ProjectConfig{})

	a.Error(err)
}
