package raindrops

import "strconv"

func Convert(number int) string {
	resultString := ""
	if number%3 == 0 {
		resultString += "Pling"
	}

	if number%5 == 0 {
		resultString += "Plang"
	}

	if number%7 == 0 {
		resultString += "Plong"
	}

	if len(resultString) == 0 {
		resultString = strconv.Itoa(number)
	}

	return resultString
}
