package thefarm

import (
	"errors"
	"fmt"
)

// See types.go for the types defined for this exercise.

// TODO: Define the SillyNephewError type here.
type SillyNephewError struct {
	cowsNumber int
}

func (n *SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", n.cowsNumber)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	availableFodder, err := weightFodder.FodderAmount()

	if err != nil {
		if err == ErrScaleMalfunction {
			if availableFodder > 0 {
				availableFodder *= 2
				return availableFodder / float64(cows), nil
			} else {
				return 0, errors.New("negative fodder")
			}

		}
		return 0, err
	}

	if availableFodder < 0 {
		return 0, errors.New("negative fodder")
	} else if cows == 0 {
		return 0, errors.New("division by zero")
	} else if cows < 0 {
		return 0, &SillyNephewError{cows}
	}

	return availableFodder / float64(cows), nil
}
