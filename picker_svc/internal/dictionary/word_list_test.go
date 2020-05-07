// Copyright 2020 Imhotep Software
// All material is licensed under the Apache License Version 2.0
// http://www.apache.org/licenses/LICENSE-2.0

package dictionary_test

import (
	"testing"

	"github.com/gopherland/labs_int/picker_svc/internal/dictionary"
	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	l, err := dictionary.LoadDefault("musicians")

	assert.Nil(t, err)
	assert.Equal(t, 5, len(l))
}

func TestLoadNoDictionary(t *testing.T) {
	_, err := dictionary.LoadDefault("actors")
	assert.EqualError(t, err, "unable to load dictionary `../assets/actors.txt")
}

func TestLoadNoAssetDir(t *testing.T) {
	_, err := dictionary.Load("zorg", "actors")
	assert.EqualError(t, err, "unable to load dictionary `zorg/actors.txt")
}
