package isogram

import (
	"strings"
	"unicode"
)

// IsIsogram returns true if input is an isogram
func IsIsogram(input string) bool {
	if input == "" {
		return true
	}
	lower := strings.ToLower(input)
	for _, c := range lower {
		if unicode.IsLetter(c) {
			if strings.Count(lower, string(c)) > 1 {
				return false
			}
		}
	}
	return true
	// normalize input: lowercase, sort, keep characters
}
