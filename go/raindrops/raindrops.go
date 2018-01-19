package raindrops

import (
	"fmt"
	"strings"
)

const testVersion = 3

// Convert converts a number into raindrop noises
func Convert(n int) string {
	return ConvertUsingJoin(n)
	// return ConvertUsingAppend(n)
}

// ConvertUsingJoin converts a number into raindrop noises using Join()
// ok  	_/Users/ferlatte/exercism/go/raindrops	2.117s
func ConvertUsingJoin(n int) string {
	var result []string
	if n%3 == 0 {
		result = append(result, "Pling")
	}
	if n%5 == 0 {
		result = append(result, "Plang")
	}
	if n%7 == 0 {
		result = append(result, "Plong")
	}
	if len(result) == 0 {
		result = append(result, fmt.Sprintf("%d", n))
	}
	return strings.Join(result, "")
}

// ConvertUsingAppend converts a number to raindrops using +=
// ok  	_/Users/ferlatte/exercism/go/raindrops	2.659s
func ConvertUsingAppend(n int) string {
	r := ""
	if n%3 == 0 {
		r += "Pling"
	}
	if n%5 == 0 {
		r += "Plang"
	}
	if n%7 == 0 {
		r += "Plong"
	}
	if len(r) == 0 {
		r = fmt.Sprintf("%d", n)
	}
	return r
}

// Don't forget the test program has a benchmark too.
// How fast does your Convert convert?
