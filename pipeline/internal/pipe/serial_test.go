// Copyright 2020 Imhotep Software
// All material is licensed under the Apache License Version 2.0
// http://www.apache.org/licenses/LICENSE-2.0

package pipe_test

import (
	"testing"

	"github.com/gopherland/labs_int/pipeline/internal/pipe"
	"github.com/stretchr/testify/assert"
)

const (
	book = "../../assets/gulliver.txt"
	word = "gulliver"
)

func TestSerial(t *testing.T) {
	count, err := pipe.Serial(book, word)
	assert.Nil(t, err)
	assert.Equal(t, 9, count)
}

func BenchmarkPipeSerial(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		pipe.Serial(book, word)
	}
}
