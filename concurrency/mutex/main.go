
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"sync"
)

type (
	bookCounts map[string]wordCounts
	wordCounts map[string]int
)

func main() {
	var (
		bc    = bookCounts{}
		books = []string{
			"100west.txt",
			"13chil.txt",
			"3gables.txt",
			"abbey.txt",
			"abyss.txt",
		}
		words = []string{
			"as",
			"the",
			"he",
			"she",
			"are",
		}
	)

	for _, book := range books {
		for _, word := range words {
			// Fire up book counter workers
			// !!YOUR_CODE!!
		}
	}
	printSummary(bc)
}

func printSummary(bc bookCounts) {
	for book, wc := range bc {
		fmt.Printf("%10s\n", book)
		for word, count := range wc {
			fmt.Printf("  %10s: %d\n", word, count)
		}
	}
}

func countWord(bc bookCounts, book, word string, wg *sync.WaitGroup, mx *sync.Mutex) {
	defer wg.Done()

	bb, err := fetch(booksURL + book)
	if err != nil {
		// Update bookCount
		// !! YOUR_CODE !!
		return
	}

	rx := regexp.MustCompile("(?i)" + word)
	count := grepCount(rx, string(bb))
	// Update bookCount
	// !! YOUR_CODE !!
}

// GrepCount returns the number of occurrence of a word in a line.
func grepCount(rx *regexp.Regexp, line string) int {
	m := rx.FindAllStringIndex(line, -1)
	return len(m)
}

// BooksURL book repo url.
const booksURL = "http://www.textfiles.com/stories/"

// Fetch a book content from a book repo.
func fetch(link string) ([]byte, error) {
	var res []byte
	resp, err := http.Get(link)
	defer func() {
		if resp != nil && resp.Body != nil {
			resp.Body.Close()
		}
	}()
	if err != nil {
		return res, err
	}
	res, err = ioutil.ReadAll(resp.Body)
	return res, err
}
