package hangman

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDictionaryLoad(t *testing.T) {
	_, err := NewDictionary("testdata/test.txt")
	assert.Nil(t, err)
}

func TestDictionaryLoadFail(t *testing.T) {
	_, err := NewDictionary("testdata/test.fred")
	assert.NotNil(t, err)
}

func TestWords(t *testing.T) {
	d, err := NewDictionary("testdata/test.txt")
	assert.Nil(t, err)
	assert.Equal(t, 5, len(d.Words()))
}
