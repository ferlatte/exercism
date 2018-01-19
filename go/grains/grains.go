package grains

import (
	"errors"
)

const testVersion = 1

// Square returns the number of grains on the square of a chessboard.
func Square(s int) (uint64, error) {
	if s < 1 || s > 64 {
		return 0, errors.New("s must be 1-64")
	}
	return 1 << uint(s-1), nil
}

// Total returns all of the grains.
func Total() uint64 {
	var total uint64
	for i := 1; i <= 64; i++ {
		s, _ := Square(i)
		total += s
	}
	return total
}
