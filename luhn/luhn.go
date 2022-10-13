package luhn

import (
	"strconv"
	"unicode"
)

func Valid(id string) bool {

	if len(id) < 2 {
		return false
	}

	digits := []int{}
	isCandidateToDouble := false

	for x := len(id) - 1; x >= 0; x-- {

		if unicode.IsDigit(rune(id[x])) {

			value, _ := strconv.Atoi(string(id[x]))

			if isCandidateToDouble {

				value *= 2

				if value > 9 {
					value -= 9
				}

				isCandidateToDouble = false
			} else {
				isCandidateToDouble = true
			}

			digits = append(digits, value)

		} else if !unicode.IsSpace(rune(id[x])) {
			return false
		}
	}

	if len(digits) < 2 {
		return false
	}

	sum := 0

	for _, value := range digits {
		sum += value
	}

	return sum%10 == 0
}
