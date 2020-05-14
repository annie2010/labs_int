package grep

// Grep returns the number of occ of a given word in a line.
func Grep(word, line string) (count int) {
	index, size := 0, len(word)
	for i := 0; i < len(line); i++ {
		if line[i] != word[index] {
			index = 0
			continue
		}
		index++
		if index == size {
			count++
			index = 0
		}
	}

	return
}
