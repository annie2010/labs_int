// Copyright 2020 Imhotep Software
// All material is licensed under the Apache License Version 2.0
// http://www.apache.org/licenses/LICENSE-2.0

package pipe

import (
	"bufio"
	"context"
	"io"
	"os"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/gopherland/labs_int/pipeline/internal"
	"github.com/gopherland/labs_int/pipeline/internal/grep"
)

type Mod int

const (
	Controlled Mod = 1 << iota
	Free
	Buffered
)

// Piped runs grep in a pipeline.
func Pipeline(ctx context.Context, mod Mod, book, word string) (int, error) {
	var (
		wg    sync.WaitGroup
		count int32
		lineC = scanner(ctx, book)
	)

	var countC chan int
	switch mod {
	case Free:
		countC = grepFree(ctx, word, lineC)
	case Controlled:
		countC = grepControlled(ctx, word, lineC)
	case Buffered:
		countC = grepBuffered(ctx, word, lineC)
	}

	for n := 0; n < 2; n++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case n, ok := <-countC:
					if !ok || n == 0 {
						return
					}
					atomic.AddInt32(&count, int32(n))
				case <-ctx.Done():
					return
				}
			}
		}()
	}

	wg.Wait()
	select {
	case err, ok := <-internal.ErrChan(ctx):
		if !ok {
			return int(count), nil
		}
		return int(count), err
	default:
		return int(count), nil
	}
}

func grepFree(ctx context.Context, word string, in chan string) chan int {
	var (
		out = make(chan int, 1)
		wg  sync.WaitGroup
	)

	go func() {
		for line := range in {
			wg.Add(1)
			go func(ctx context.Context, line string, out chan<- int) {
				defer wg.Done()
				count := grep.Grep(word, line)
				if count == 0 {
					return
				}
				select {
				case out <- count:
				case <-ctx.Done():
					return
				}
			}(ctx, line, out)
		}
		go func() {
			wg.Wait()
			close(out)
		}()
	}()

	return out
}

func grepBuffered(ctx context.Context, word string, in chan string) chan int {
	var (
		out = make(chan int, 1)
		wg  sync.WaitGroup
		sem = make(chan struct{}, 10)
	)

	go func() {
		for line := range in {
			sem <- struct{}{}
			wg.Add(1)
			go func(ctx context.Context, sem <-chan struct{}, line string, out chan<- int) {
				defer func() {
					<-sem
					wg.Done()
				}()
				count := grep.Grep(word, line)
				if count == 0 {
					return
				}
				select {
				case out <- count:
				case <-ctx.Done():
					return
				}
			}(ctx, sem, line, out)
		}

		go func() {
			wg.Wait()
			close(out)
		}()
	}()

	return out
}

func grepControlled(ctx context.Context, word string, in chan string) chan int {
	var (
		out = make(chan int, 3)
		wg  sync.WaitGroup
	)
	for n := 0; n < 3; n++ {
		wg.Add(1)
		go func(ctx context.Context, in <-chan string, out chan<- int) {
			defer wg.Done()

			for line := range in {
				count := grep.Grep(word, line)
				if count == 0 {
					continue
				}
				select {
				case out <- count:
				case <-ctx.Done():
					return
				}
			}
		}(ctx, in, out)
	}
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func scanner(ctx context.Context, path string) chan string {
	out := make(chan string, 10)

	go func(ctx context.Context, out chan<- string) {
		defer func() {
			close(out)
		}()
		file, err := os.Open(path)
		if err != nil {
			internal.ErrChan(ctx) <- err
			return
		}
		defer func() {
			_ = file.Close()
		}()

		buff := bufio.NewReader(file)
		for {
			line, err := buff.ReadString('\n')
			if err == io.EOF {
				line = strings.ToLower(strings.TrimSpace(line))
				if len(line) > 0 {
					out <- line
				}
				break
			}
			if err != nil {
				internal.ErrChan(ctx) <- err
				return
			}
			line = strings.ToLower(strings.TrimSpace(line))
			if len(line) > 0 {
				out <- line
			}
		}
	}(ctx, out)

	return out
}
