package diffsquares

const testVersion = 1

// SumOfSquares returns the sum of the squares of the first n natural numbers.
func SumOfSquares(n int) int {
	return n * (n + 1) * (2*n + 1) / 6
}

// SquareOfSums returns the square of the sums of the first n natural numbers.
func SquareOfSums(n int) int {
	sum := n * (1 + n) / 2
	return sum * sum
}

// Difference returns the difference between the square of the sum and the sum
// of the squares of the first n natural numbers.
func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}
