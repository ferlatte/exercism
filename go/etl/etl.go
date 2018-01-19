package etl

import (
	"strings"
)

const testVersion = 1

// Transform Scrabble score data from score -> array of letters to
// letter -> score
func Transform(m map[int][]string) map[string]int {
	retval := make(map[string]int)
	for score, letters := range m {
		for _, l := range letters {
			lc := strings.ToLower(l)
			retval[lc] = score
		}
	}
	return retval
}
