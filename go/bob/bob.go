package bob

import (
	"regexp"
	"unicode"
)

const testVersion = 3

// Hey returns what bob says in response to what you say.
func Hey(s string) string {
	//
	if isEmpty(s) {
		return "Fine. Be that way!"
	}
	if isShouty(s) {
		return "Whoa, chill out!"
	}
	if isQuestion(s) {
		return "Sure."
	}
	return "Whatever."
}

// Ensure that all letters in the string are uppercase
func isShouty(s string) bool {
	atLeastOneLetter := false
	for _, r := range s {
		if unicode.IsLetter(r) {
			atLeastOneLetter = true
			if unicode.IsLower(r) {
				return false
			}
		}
	}
	// true means we saw at least one letter AND they were all Upper
	return atLeastOneLetter
}

func isQuestion(s string) bool {
	pattern := `\?[[:blank:]]*$`
	return match(pattern, s)
}

func isEmpty(s string) bool {
	pattern := `^\s*$`
	return match(pattern, s)
}

func match(pattern, s string) bool {
	matched, err := regexp.MatchString(pattern, s)
	if err != nil {
		panic(err)
	}
	return matched
}
