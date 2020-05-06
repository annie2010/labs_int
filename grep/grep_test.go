// Copyright 2020 Imhotep Software
// All material is licensed under the Apache License Version 2.0
// http://www.apache.org/licenses/LICENSE-2.0

package grep_test

import (
	"testing"

	"github.com/gopherland/labs_int/grep"
	"gotest.tools/assert"
)

func TestWordCount(t *testing.T) {
	uu := map[string]struct {
		text string
		e    int64
	}{
		"semi-cols": {
			text: "*** START OF THIS PROJECT GUTENBERG EBOOK MOBY DICK; OR THE WHALE ***",
			e:    1,
		},
		"dash": {
			text: "MOBY-DICK;",
			e:    1,
		},
		"quote": {
			text: `“Moby Dick?” shouted Ahab. “Do ye know the white whale then, Tash?”`,
			e:    1,
		},
		"special": {
			text: "seen—Moby Dick—Moby Dick!”",
			e:    2,
		},
	}

	t.Parallel()
	for k := range uu {
		u := uu[k]
		t.Run(k, func(t *testing.T) {
			assert.Equal(t, u.e, grep.WordCount("moby", u.text))
		})
	}
}

func TestWordCountBytes(t *testing.T) {
	uu := map[string]struct {
		text string
		e    int64
	}{
		"semi-cols": {
			text: "*** START OF THIS PROJECT GUTENBERG EBOOK MOBY DICK; OR THE WHALE ***",
			e:    1,
		},
		"dash": {
			text: "MOBY-DICK;",
			e:    1,
		},
		"quotes": {
			text: `“Moby Dick?” shouted Ahab. “Do ye know the white whale then, Tash?”`,
			e:    1,
		},
		"special-dash": {
			text: "seen—Moby Dick—Moby Dick!”",
			e:    2,
		},
	}

	t.Parallel()
	for k := range uu {
		u := uu[k]
		t.Run(k, func(t *testing.T) {
			assert.Equal(t, u.e, grep.WordCountBytes("moby", u.text))
		})
	}
}

func BenchmarkWordCountV1(b *testing.B) {
	const text = `Moby Dick?” shouted Ahab. “Do ye know the white whale then, Tash?`

	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		grep.WordCount("moby", text)
	}
}

func BenchmarkWordCountV2(b *testing.B) {
	const text = `Moby Dick?” shouted Ahab. “Do ye know the white whale then, Tash?`

	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		grep.WordCountBytes("moby", text)
	}
}
