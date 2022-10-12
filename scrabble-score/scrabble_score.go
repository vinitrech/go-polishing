package scrabble

import (
	"strings"
)

func Score(word string) int {
	counter := 0

	for _, letter := range word {
		toUpper := []rune(strings.ToUpper(string(letter)))

		counter += getValue(toUpper[0])
	}

	return counter
}

func getValue(letter rune) int {
	switch letter {
	case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
		return 1
	case 'D', 'G':
		return 2
	case 'B', 'C', 'M', 'P':
		return 3
	case 'F', 'H', 'V', 'W', 'Y':
		return 4
	case 'K':
		return 5
	case 'J', 'X':
		return 8
	case 'Q', 'Z':
		return 10
	}
	return 0
}
