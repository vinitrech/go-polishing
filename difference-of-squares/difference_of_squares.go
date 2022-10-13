package diffsquares

func SquareOfSum(n int) int {
	sum := 0

	for x := 0; x <= n; x++ {
		sum += x
	}

	return sum * sum
}

func SumOfSquares(n int) int {
	sum := 0

	for x := 0; x <= n; x++ {
		sum += x * x
	}

	return sum
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
