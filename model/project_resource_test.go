package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProjectResource_PathFor(t *testing.T) {
	a := assert.New(t)
	r := ProjectResource{
		PathExpression: `{{ .Dir }}/{{ .Locale | replace "-" "_" }}{{ .Ext }}`,
	}
	p, err := r.PathFor("testdata/en-US.json", "en-US")

	a.NoError(err)
	a.EqualValues("testdata/en_US.json", p)
}

func TestProjectResource_PathFor_ThrowsError(t *testing.T) {
	a := assert.New(t)
	r := ProjectResource{
		PathExpression: "{{ .Test | test }}",
	}
	p, err := r.PathFor("testdata/en-US.json", "en-US")

	a.Error(err)
	a.EqualValues("", p)
}

func TestProjectResource_Files(t *testing.T) {
	a := assert.New(t)
	r1 := ProjectResource{
		PathGlob: "testdata/**/en-US.json",
	}
	r2 := ProjectResource{
		PathGlob: "testdata/**/en-US.json",
		PathExclude: []string{
			"testdata/foo/en-US.json",
		},
	}
	r3 := ProjectResource{
		PathGlob: "testdata/**/en-US.json",
		PathExclude: []string{
			"testdata/foo/en-US.json",
			"testdata/bar/en-US.json",
		},
	}
	r4 := ProjectResource{
		PathGlob: "testdata/**/en-US.json",
		PathExclude: []string{
			"testdata/foo/.*",
		},
	}

	a.EqualValues([]string{
		"testdata/bar/en-US.json",
		"testdata/foo/en-US.json",
	}, r1.Files())
	a.EqualValues([]string{
		"testdata/bar/en-US.json",
	}, r2.Files())
	a.EqualValues([]string{}, r3.Files())
	a.EqualValues([]string{
		"testdata/bar/en-US.json",
	}, r4.Files())
}
