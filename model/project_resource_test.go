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

func TestProjectResource_LimitFiles(t *testing.T) {
	a := assert.New(t)
	r := ProjectResource{
		PathGlob: "testdata/**/*.json",
	}

	f1 := r.LimitFiles(-1)
	f2 := r.LimitFiles(1)
	f3 := r.LimitFiles(3)
	f4 := r.LimitFiles(4)

	a.Len(f1, 1)
	a.Len(f1[0], 7)

	a.Len(f2, 7)
	a.Len(f2[0], 1)
	a.Len(f2[1], 1)
	a.Len(f2[2], 1)
	a.Len(f2[3], 1)
	a.Len(f2[4], 1)
	a.Len(f2[5], 1)
	a.Len(f2[6], 1)

	a.Len(f3, 3)
	a.Len(f3[0], 3)
	a.Len(f3[1], 3)
	a.Len(f3[2], 1)

	a.Len(f4, 2)
	a.Len(f4[0], 4)
	a.Len(f4[1], 3)
}
