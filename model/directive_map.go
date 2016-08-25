package model

import (
	"strings"
)

const DirectivePrefix = "smartling."

type DirectiveMap map[string]string

func (d DirectiveMap) WithPrefix() map[string]string {
	m := map[string]string{}

	for key, value := range d {
		if !strings.HasPrefix(key, DirectivePrefix) {
			m[DirectivePrefix+key] = value
		} else {
			m[key] = value
		}
	}

	return m
}
