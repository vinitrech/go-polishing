package luhn

import "fmt"

func Valid(id string) bool {

	if len(id) < 2 {
		return false
	}

	count, sum := 0, 0

	for x := len(id) - 1; x >= 0; x-- {

		if id[x] >= '0' && id[x] <= '9' {

			value := int(id[x] - '0')

			if count%2 == 1 {

				value <<= 1

				fmt.Println("Value: ", value, " - id[x]", id[x], " - rune", '0')

				if value > 9 {
					value -= 9
				}

			}

			sum += value
			count++

		} else if id[x] != ' ' {
			return false
		}
	}

	return sum%10 == 0 && count > 1
}
