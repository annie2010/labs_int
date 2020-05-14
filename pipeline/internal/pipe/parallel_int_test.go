// Copyright 2020 Imhotep Software
// All material is licensed under the Apache License Version 2.0
// http://www.apache.org/licenses/LICENSE-2.0

package pipe

import (
	"testing"

	"github.com/gopherland/labs_int/pipeline/internal"
	"github.com/stretchr/testify/assert"
)

func TestGrepFree(t *testing.T) {
	ctx := internal.NewContext()
	in := make(chan string, 100)

	c := grepFree(ctx, "blah", in)
	for i := 0; i < 100; i++ {
		in <- "blah blah gulliver"
	}
	close(in)
	var count int
	for n := range c {
		count += n
	}
	assert.Equal(t, 200, count)
}

func TestGrepControlled(t *testing.T) {
	ctx := internal.NewContext()
	in := make(chan string, 100)

	c := grepControlled(ctx, "blah", in)
	for i := 0; i < 100; i++ {
		in <- "blah blah gulliver"
	}
	close(in)
	var count int
	for n := range c {
		count += n
	}
	assert.Equal(t, 200, count)
}

func TestGrepBuffered(t *testing.T) {
	ctx := internal.NewContext()
	in := make(chan string, 100)

	c := grepBuffered(ctx, "blah", in)
	for i := 0; i < 100; i++ {
		in <- "blah blah gulliver"
	}
	close(in)
	var count int
	for n := range c {
		count += n
	}
	assert.Equal(t, 200, count)
}

func BenchmarkGrepFree(b *testing.B) {
	ctx := internal.NewContext()

	in := make(chan string)
	c := grepFree(ctx, "blah", in)
	for i := 0; i < 1; i++ {
		go func() {
			for range c {
			}
		}()
	}
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		for i := 0; i < 200; i++ {
			in <- "blah blah gulliver"
		}
	}
	close(in)
}

func BenchmarkGrepControlled(b *testing.B) {
	ctx := internal.NewContext()

	in := make(chan string)
	c := grepControlled(ctx, "blah", in)
	for i := 0; i < 1; i++ {
		go func() {
			for range c {
			}
		}()
	}
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		for i := 0; i < 200; i++ {
			in <- "blah blah gulliver"
		}
	}
	close(in)
}

func BenchmarkGrepBuffered(b *testing.B) {
	ctx := internal.NewContext()

	in := make(chan string)
	c := grepBuffered(ctx, "blah", in)
	for i := 0; i < 1; i++ {
		go func() {
			for range c {
			}
		}()
	}
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		for i := 0; i < 200; i++ {
			in <- "blah blah gulliver"
		}
	}
	close(in)
}
