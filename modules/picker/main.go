// Copyright 2020 Imhotep Software
// All material is licensed under the Apache License Version 2.0
// http://www.apache.org/licenses/LICENSE-2.0

package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/gookit/color"
	"github.com/gopherland/labs_int/modules/dictionary"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	path := flag.String("p", "assets/words.txt", "Specifies a dictionary path")
	flag.Parse()

	wl, err := dictionary.Load(*path, dictionary.WordList{})
	if err != nil {
		panic(err)
	}
	if len(wl) == 0 {
		panic("No words found in dictionary")
	}

	fmt.Print("\033[H\033[2J")
	cyan := color.FgCyan.Render
	fmt.Printf("🧠  >>> Word of the day is %s <<<  🤡\n", cyan(pick(wl)))
}

func pick(l dictionary.WordList) string {
	return strings.TrimSpace(l[rand.Intn(len(l))])
}
