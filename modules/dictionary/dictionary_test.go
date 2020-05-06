// Copyright 2020 Imhotep Software
// All material is licensed under the Apache License Version 2.0
// http://www.apache.org/licenses/LICENSE-2.0

package dictionary_test

import (
	"testing"

	"github.com/gopherland/labs_int/modules/dictionary"
	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	uu := map[string]struct {
		path     string
		excludes []string
		wl       dictionary.WordList
	}{
		"no-excludes": {
			path: "testdata/dic1.txt",
			wl:   dictionary.WordList{"fred", "blee", "duh"},
		},
		"excludes": {
			path:     "testdata/dic1.txt",
			excludes: dictionary.WordList{"fred", "blee"},
			wl:       dictionary.WordList{"duh"},
		},
		"missing": {
			path:     "testdata/dic1.txt",
			excludes: dictionary.WordList{"fred", "blee1", "zorg"},
			wl:       dictionary.WordList{"blee", "duh"},
		},
	}

	for k := range uu {
		u := uu[k]
		t.Run(k, func(t *testing.T) {
			wl, err := dictionary.Load(u.path, u.excludes)
			assert.Nil(t, err)
			assert.Equal(t, u.wl, wl)
		})
	}
}
