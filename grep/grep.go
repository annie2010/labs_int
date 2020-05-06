// © 2020 Imhotep Software LLC. All rights reserved.

package grep

import (
	"regexp"
	"strings"
)

// Count returns the number of occurrence of a word in a line.
func Count(word, line string) int64 {
	var rx = regexp.MustCompile(`[,.\-_,;“—‘]`)
	l := strings.ToLower(line)
	l = rx.ReplaceAllString(l, " ")

	var count int64
	tokens := strings.Split(l, " ")
	for _, t := range tokens {
		if strings.TrimSpace(t) == word {
			count++
		}
	}

	return count
}

// Count1 returns the number of occurrence of a word in a line.
func Count1(word, line string) int64 {
	l := strings.ToLower(line)

	var index int
	var count int64
	for _, b := range []byte(l) {
		if b != word[index] {
			index = 0
			continue
		}

		index++
		if index == len(word) {
			count++
			index = 0
		}
	}

	return count
}
