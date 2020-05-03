// © 2020 Imhotep Software LLC. All rights reserved.

package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/gopherland/level2/labs/grep"
)

func main() {
	var word, fileName string
	flag.StringVar(&word, "w", "moby", "Specify the word to find")
	flag.StringVar(&fileName, "f", "assets/mobby.txt", "Specify a text file to grep from")
	flag.Parse()

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Unable to open file %s", fileName)
	}

	scanner := bufio.NewScanner(file)
	var count int

	for scanner.Scan() {
		count += grep.Count(word, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Scanner failed %v", err)
	}

	fmt.Printf("Found %d occurrences of %q\n", count, word)
}
