// Copyright 2020 Imhotep Software
// All material is licensed under the Apache License Version 2.0
// http://www.apache.org/licenses/LICENSE-2.0

package pipe

import (
	"bufio"
	"io"
	"os"
	"strings"

	"github.com/gopherland/labs_int/pipeline/internal/grep"
)

// Serial runs grep in serial.
func Serial(book, word string) (int, error) {
	file, err := os.Open(book)
	if err != nil {
		return 0, err
	}
	defer func() {
		_ = file.Close()
	}()

	var count int
	buff := bufio.NewReader(file)
	for {
		line, err := buff.ReadString('\n')
		if err == io.EOF {
			count += grep.Grep(word, strings.ToLower(strings.TrimSpace(line)))
			break
		}
		if err != nil {
			return 0, err
		}
		count += grep.Grep(word, strings.ToLower(strings.TrimSpace(line)))
	}

	return count, nil
}
