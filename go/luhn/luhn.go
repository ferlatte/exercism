package luhn

import "unicode"

var evens = map[int]int{
	0: 0,
	1: 2,
	2: 4,
	3: 6,
	4: 8,
	5: 1,
	6: 3,
	7: 5,
	8: 7,
	9: 9,
}

// Valid tells you if a string has a valid Luhn checksum.
func Valid(input string) bool {
	even := false
	total := 0
	runes := []rune(input)
	n := 0
	for i := len(runes) - 1; i >= 0; i-- {
		r := runes[i]
		if unicode.IsSpace(r) {
			continue
		}
		if !unicode.IsNumber(r) {
			return false
		}
		v := int(r - '0')
		if even {
			v = evens[v]
			even = false
		} else {
			even = true
		}
		n++
		total += v
	}
	if n < 2 {
		return false
	}
	return total%10 == 0
}
