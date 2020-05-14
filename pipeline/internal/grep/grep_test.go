// Copyright 2020 Imhotep Software
// All material is licensed under the Apache License Version 2.0
// http://www.apache.org/licenses/LICENSE-2.0

package grep_test

import (
	"strings"
	"testing"

	"github.com/gopherland/labs_int/pipeline/internal/grep"
	"github.com/stretchr/testify/assert"
)

func TestGrep(t *testing.T) {
	samples := genSamples(t)

	uu := map[string]struct {
		text string
		e    int
	}{
		"semi-cols": {
			text: samples[0],
			e:    1,
		},
		"dash": {
			text: samples[1],
			e:    1,
		},
		"quotes": {
			text: samples[2],
			e:    1,
		},
		"special-dash": {
			text: samples[3],
			e:    2,
		},
	}

	t.Parallel()
	for k := range uu {
		u := uu[k]
		t.Run(k, func(t *testing.T) {
			assert.Equal(t, u.e, grep.Grep("moby", strings.ToLower(u.text)))
		})
	}
}

func BenchmarkGrep(b *testing.B) {
	sample := genSamples(b)[2]

	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		grep.Grep("moby", sample)
	}
}

// Helpers...

func genSamples(t testing.TB) []string {
	t.Helper()

	return []string{
		"*** START OF THIS PROJECT GUTENBERG EBOOK MOBY DICK; OR THIS WHALE ***",
		"MOBY-DICK;",
		`“Moby Dick?” shouted Ahab. “Do ye know the white whale then, Tash?”`,
		"seen—Moby Dick—Moby Dick!”",
	}
}
