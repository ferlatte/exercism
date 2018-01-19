package hamming

import (
	"errors"
)

const testVersion = 6

// Distance computes the Hamming distance of two DNA strings
// Surprisingly to me, DistanceUnicode is almost a second faster
// according to go test -bench .
func Distance(a, b string) (int, error) {
	// return DistanceASCII(a, b)
	return DistanceUnicode(a, b)
}

// DistanceASCII computes the Hamming distance of two strings, ASCII only
// Benchmark: ok  	_/Users/ferlatte/exercism/go/hamming	2.475s
func DistanceASCII(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("a,b must be same length")
	}

	distance := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			distance++
		}
	}
	return distance, nil
}

// DistanceUnicode computes the Hamming distance of two strings, Unicode
// Benchmark: ok  	_/Users/ferlatte/exercism/go/hamming	1.845s
func DistanceUnicode(a, b string) (int, error) {
	aRunes := []rune(a)
	bRunes := []rune(b)
	distance := 0
	if len(aRunes) != len(bRunes) {
		return 0, errors.New("a,b must be same length")
	}
	for i := 0; i < len(aRunes); i++ {
		if aRunes[i] != bRunes[i] {
			distance++
		}
	}
	return distance, nil
}
