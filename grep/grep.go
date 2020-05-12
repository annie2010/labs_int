// Copyright 2020 Imhotep Software
// All material is licensed under the Apache License Version 2.0
// http://www.apache.org/licenses/LICENSE-2.0

package grep

import (
	"regexp"
	"strings"
)

// WordCount returns the number of occurrence of a word in a line.
func WordCount(word, line string) (count int64) {
	var (
		rx = regexp.MustCompile(`[,.\-_,;“—‘]`)
		l  = rx.ReplaceAllString(strings.ToLower(line), " ")
	)
	tokens := strings.Split(l, " ")
	for _, t := range tokens {
		if strings.TrimSpace(t) == word {
			count++
		}
	}

	return
}

// WordCountBytes returns the number of occurrence of a word in a line.
func WordCountBytes(word, line string) (count int64) {
	var (
		index int
		l     = strings.ToLower(line)
	)
	for _, b := range []byte(l) {
		<<!!YOUR_CODE!!>> Iterate thru the line a byte at a time and compare with the byte in the word
	}

	return
}
