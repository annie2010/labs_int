// Copyright 2020 Imhotep Software
// All material is licensed under the Apache License Version 2.0
// http://www.apache.org/licenses/LICENSE-2.0

// Package dictionary represents a words dictionary.
package dictionary

import (
	"github.com/gopherland/labs_int/modules/dictionary/internal"
)

// WordList a collection of dictionary words.
type WordList []string

// Load reads a collection of words from a given a path.
func Load(path string, excludes WordList) (WordList, error) {
	<<!!YOUR_CODE!!>> -- Leveraging internal.Load load your word list
	<<!!YOUR_CODE!!>> -- Implement the words exclusions functionality. Make sure all tests pass!
	<<!!YOUR_CODE!!>> -- BONUS - Implement in place exclusions!
}
