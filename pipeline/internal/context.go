// Copyright 2020 Imhotep Software
// All material is licensed under the Apache License Version 2.0
// http://www.apache.org/licenses/LICENSE-2.0

package internal

import "context"

type contextKey string

const errKey contextKey = "errChan"

// NewContext returns a new pipeline context.
func NewContext() context.Context {
	return context.WithValue(
		context.Background(),
		errKey,
		make(chan error, 1),
	)
}

func ErrChan(ctx context.Context) chan error {
	e, ok := ctx.Value(errKey).(chan error)
	if !ok {
		return nil
	}
	return e
}
