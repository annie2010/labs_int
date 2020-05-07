// Copyright 2020 Imhotep Software
// All material is licensed under the Apache License Version 2.0
// http://www.apache.org/licenses/LICENSE-2.0

// Package dictionary picks a random word from a dictionary loaded from
// that can be randomly picked.
package dictionary

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

const defaultAssetDir = "../assets"

// WordList a collection of dictionary words
type WordList []string

// LoadDefault loads a dictionary by name from the default assets dir
func LoadDefault(name string) (wl WordList, e error) {
	return Load(defaultAssetDir, name)
}

// Load loads a new dictionary given a base path and a dictionary name
func Load(dir, name string) (wl WordList, e error) {
	path := filepath.Join(dir, name+".txt")

	if f, err := os.Open(path); err != nil {
		e = fmt.Errorf("unable to load dictionary `%s", path)
	} else {
		sc := bufio.NewScanner(f)
		for sc.Scan() {
			wl = append(wl, sc.Text())
		}
	}
	return
}
