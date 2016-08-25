package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProjectConfig_Merge(t *testing.T) {
	a := assert.New(t)
	src := ProjectConfig{
		AuthToken: AuthToken{
			AccessToken: "accessToken1",
		},
		UserToken: UserToken{
			ID: "id1",
		},
		Project: Project{
			Alias: "projectAlias1",
		},
	}
	dst := ProjectConfig{
		AuthToken: AuthToken{
			AccessToken: "accessToken2",
		},
		UserToken: UserToken{
			ID: "id2",
		},
		Project: Project{
			Alias: "projectAlias2",
		},
	}

	err := src.Merge(&dst)

	a.NoError(err)
	a.EqualValues(dst, src)
}

func TestProjectConfig_LocaleFor(t *testing.T) {
	a := assert.New(t)
	p := ProjectConfig{
		Locales: map[string]string{
			"en-US": "en_US",
			"ru-RU": "ru",
		},
	}

	a.EqualValues("en_US", p.LocaleFor("en-US"))
	a.EqualValues("ru", p.LocaleFor("ru-RU"))
	a.EqualValues("de-DE", p.LocaleFor("de-DE"))
}

func TestProjectConfig_FileURI(t *testing.T) {
	a := assert.New(t)
	p1 := ProjectConfig{
		Project: Project{
			Alias: "testdata",
		},
	}
	p2 := ProjectConfig{}

	a.EqualValues("testdata/foo.json", p1.FileURI("foo.json"))
	a.EqualValues("foo.json", p2.FileURI("foo.json"))
}

func TestProjectConfig_FilePath(t *testing.T) {
	a := assert.New(t)
	p1 := ProjectConfig{
		Project: Project{
			Alias: "testdata",
		},
	}
	p2 := ProjectConfig{}

	a.EqualValues("foo.json", p1.FilePath("testdata/foo.json"))
	a.EqualValues("foo.json", p2.FilePath("foo.json"))
	a.EqualValues("testdata/foo.json", p2.FilePath("testdata/foo.json"))
}
