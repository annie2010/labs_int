// Copyright 2020 Imhotep Software
// All material is licensed under the Apache License Version 2.0
// http://www.apache.org/licenses/LICENSE-2.0

package picker_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	l, err := picker.LoadDefault("musicians")

	assert.Nil(t, err)
	assert.Equal(t, 5, len(l))
}

func TestPicker(t *testing.T) {
	l, err := picker.LoadDefault("musicians")
	w := picker.Pick(l)

	assert.Nil(t, err)
	assert.NotEqual(t, "", w)
}

func TestLoadMissing(t *testing.T) {
	_, err := picker.LoadDefault("actors")
	assert.EqualError(t, err, "unable to load dictionary `assets/actors.txt")
}

func TestLoadCustomFail(t *testing.T) {
	_, err := picker.Load("zorg", "actors")
	assert.EqualError(t, err, "unable to load dictionary `zorg/actors.txt")
}
