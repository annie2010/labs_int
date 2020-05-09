// Copyright 2020 Imhotep Software
// All material is licensed under the Apache License Version 2.0
// http://www.apache.org/licenses/LICENSE-2.0

package magnifica_test

import (
	"encoding/json"
	"testing"

	"github.com/gopherland/labs_int/magnifica"
	"github.com/stretchr/testify/assert"
)

func TestEntryMarshal(t *testing.T) {
	e := magnifica.Entry{
		Dictionary: "slang",
		Location:   "/acme/english",
		Word:       "bumblebeetuna",
	}

	bb, err := json.Marshal(&e)
	assert.Nil(t, err)
	assert.Equal(t, string(raw), string(bb))
}

func TestEntryUnmarshal(t *testing.T) {
	var e magnifica.Entry
	err := json.Unmarshal([]byte(raw), &e)

	assert.Nil(t, err)
	assert.Equal(t, "slang", e.Dictionary)
	assert.Equal(t, "/acme/english/", e.Location)
	assert.Equal(t, "bumblebeetuna", e.Word)
}

const raw = `{"dictionary_location":"/acme/english/slang","dictionary_word":"bumblebeetuna","political_correctness":false}`
