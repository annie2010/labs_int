// Copyright 2020 Imhotep Software
// All material is licensed under the Apache License Version 2.0
// http://www.apache.org/licenses/LICENSE-2.0

package main

import (
	"flag"
	"fmt"
)

func main() {
	var (
		dic = flag.String("dic", "default", "Specifies a dictionary name")
		dir = flag.String("dir", "./assets", "Specifies dictionaries load directory")
	)
	flag.Parse()

	if wl, err := picker.Load(*dir, *dic); err != nil {
		panic(err)
	} else {
		fmt.Printf("The word of the day is `%s\n", picker.Pick(wl))
	}
}
