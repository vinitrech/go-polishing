package hamming

import "fmt"

func Distance(a, b string) (int, error) {
	counter := 0

	if len(a) != len(b) {
		return 0, fmt.Errorf("different lengths")
	}

	for index, cell := range a {
		if rune(b[index]) != cell {
			counter++
		}
	}

	return counter, nil
}
