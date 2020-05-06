// Copyright 2020 Imhotep Software
// All material is licensed under the Apache License Version 2.0
// http://www.apache.org/licenses/LICENSE-2.0

// Copyright 2020 Imhotep Software
// All material is licensed under the Apache License Version 2.0
// http://www.apache.org/licenses/LICENSE-2.0

package internal_test

import (
	"testing"

	"github.com/gopherland/labs_int/modules/dictionary/internal"
	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	wl, err := internal.Load("../testdata/dic1.txt")
	assert.Nil(t, err)
	assert.Equal(t, 3, len(wl))
}

func TestLoadMissing(t *testing.T) {
	_, err := internal.Load("../testdata/fred.txt")
	assert.EqualError(t, err, "unable to load dictionary `../testdata/fred.txt")
}
