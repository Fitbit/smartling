package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDirectiveMap_WithPrefix(t *testing.T) {
	a := assert.New(t)
	d := DirectiveMap{
		"string_format":          "NONE",
		"smartling.file_charset": "UTF-16",
	}
	m := map[string]string{
		"smartling.string_format": "NONE",
		"smartling.file_charset":  "UTF-16",
	}

	a.EqualValues(m, d.WithPrefix())
}
