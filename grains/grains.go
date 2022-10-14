package grains

import "fmt"

func Square(number int) (uint64, error) {

	if number > 0 && number <= 64 {
		return 1 << (number - 1), nil
	}

	return 0, fmt.Errorf("")
}

func Total() uint64 {
	return 1<<64 - 1
}
