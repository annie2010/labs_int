package hangman

import (
	"bufio"
	"os"
	"path"
)

type (
	// WordList a collection of strings
	WordList []string

	// Dictionary tracks all available words
	Dictionary struct {
		words WordList
	}
)

// NewDictionary creates a new dictionary
func NewDictionary(dir, dic string) (*Dictionary, error) {
	wl, err := load(dir, dic)
	if err != nil {
		return nil, err
	}
	return &Dictionary{words: wl}, nil
}

// Words list all loaded dictionary words
func (d *Dictionary) Words() []string {
	return d.words
}

func load(dir, dic string) (WordList, error) {
	var wl WordList
	file := path.Join(dir, dic)

	f, err := os.Open(file)
	if err != nil {
		return wl, err
	}
	s := bufio.NewScanner(f)
	for s.Scan() {
		wl = append(wl, s.Text())
	}
	return wl, err
}
