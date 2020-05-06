package internal

import(
	"bufio"
	"os"
	"fmt"
)

// Load loads a collection of words from a given a path.
func Load(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("unable to load dictionary `%s", path)
	}

	wl := make([]string, 0, 100)
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		wl = append(wl, sc.Text())
	}
	return wl, nil
}
