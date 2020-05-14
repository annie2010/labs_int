// Copyright 2020 Imhotep Software
// All material is licensed under the Apache License Version 2.0
// http://www.apache.org/licenses/LICENSE-2.0

package pipe_test

import (
	"context"
	"testing"

	"github.com/gopherland/labs_int/pipeline/internal/pipe"
	"github.com/stretchr/testify/assert"
)

func TestPipeline(t *testing.T) {
	count, err := pipe.Pipeline(context.Background(), pipe.Controlled, book, word)
	assert.Nil(t, err)
	assert.Equal(t, 9, count)
}
