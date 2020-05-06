// Copyright 2020 Imhotep Software
// All material is licensed under the Apache License Version 2.0
// http://www.apache.org/licenses/LICENSE-2.0

package internal

import (
	"bufio"
	"fmt"
	"os"
)

// Load loads a collection of words from a given a path.
func Load(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("unable to load dictionary `%s", path)
	}

	wl := make([]string, 0, 100)
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		wl = append(wl, sc.Text())
	}
	return wl, nil
}
