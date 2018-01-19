package triangle

import (
	"math"
)

const testVersion = 3

// Kind is the type of triangle
type Kind int

// Pick values for the following identifiers used by the test program.
const (
	NaT Kind = iota // not a triangle
	Equ             // equilateral
	Iso             // isosceles
	Sca             // scalene
)

// KindFromSides returns the Kind of a triangle represented by 3 sides.
func KindFromSides(a, b, c float64) Kind {
	if !isValidTriangle(a, b, c) {
		return NaT
	}
	if a == b && b == c {
		return Equ
	}
	if a == b || b == c || a == c {
		return Iso
	}
	return Sca
}

func isValidSide(s float64) bool {
	return s > 0 && !math.IsInf(s, 0) && !math.IsNaN(s)
}

func isValidTriangle(a, b, c float64) bool {
	if !isValidSide(a) || !isValidSide(b) || !isValidSide(c) {
		return false
	}
	if a+b >= c && a+c >= b && b+c >= a {
		return true
	}
	return false
}
