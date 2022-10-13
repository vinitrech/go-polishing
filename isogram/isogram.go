package isogram

import (
	"unicode"
)

func IsIsogram(word string) bool {
	repeated := map[rune]int{}

	for _, character := range word {
		if unicode.IsLetter(character) {
			toLower := unicode.ToLower(character)

			if repeated[toLower] > 0 {
				return false
			}

			repeated[toLower] += 1
		}
	}

	return true
}
