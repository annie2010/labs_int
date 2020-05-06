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
	wl, err := internal.Load(path)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(wl); i++ {
		if in(excludes, wl[i]) {
			wl = append(wl[:i], wl[i+1:]...)
			i--
		}
	}
	return wl, nil
}

func in(ww WordList, word string) bool {
	for _, w := range ww {
		if w == word {
			return true
		}
	}

	return false
}
