package interest

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	var interest float32

	if balance < 0 {
		interest = 3.213
	} else if balance >= 0 && balance < 1000 {
		interest = 0.5
	} else if balance >= 1000 && balance < 5000 {
		interest = 1.621
	} else {
		interest = 2.475
	}

	return interest
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	var result float64

	interest := InterestRate(balance)

	result = (balance * float64(interest)) / 100

	return result
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	return balance + Interest(balance)
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	years := 0

	for balance < targetBalance {
		balance += Interest(balance)
		years++
	}

	return years
}
