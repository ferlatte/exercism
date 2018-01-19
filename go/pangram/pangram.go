package pangram

import (
	"strings"
)

const testVersion = 1

// IsPangram returns true if s is a pangram.
func IsPangram(s string) bool {
	m := make(map[string]bool)
	m["a"] = false
	m["b"] = false
	m["c"] = false
	m["d"] = false
	m["e"] = false
	m["f"] = false
	m["g"] = false
	m["h"] = false
	m["i"] = false
	m["j"] = false
	m["k"] = false
	m["l"] = false
	m["m"] = false
	m["n"] = false
	m["o"] = false
	m["p"] = false
	m["q"] = false
	m["r"] = false
	m["s"] = false
	m["t"] = false
	m["u"] = false
	m["v"] = false
	m["w"] = false
	m["x"] = false
	m["y"] = false
	m["z"] = false

	lowerCased := strings.ToLower(s)
	for _, r := range lowerCased {
		key := string(r)
		_, ok := m[key]
		if ok {
			m[key] = true
		}
	}
	for _, value := range m {
		if !value {
			return false
		}
	}
	return true
}
