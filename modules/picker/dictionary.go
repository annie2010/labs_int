// © 2020 Imhotep Software LLC. All rights reserved.

// Package picker represents a random word picker from a dictionary.
package picker

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

const rootDir = "./assets"

// WordList a collection of dictionary words.
type WordList []string

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Load reads a collection of words from a given a path and a dictionary name.
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

// LoadDefault loads a dictionary by name from the default assets dir
func LoadDefault(name string) (wl WordList, e error) {
	return Load(rootDir, name)
}

// Pick select a word at random from a list of words
func Pick(l WordList) string {
	return l[rand.Intn(len(l))]
}
