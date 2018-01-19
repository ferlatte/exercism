package series

const testVersion = 2

// All returns all series of length n from string s
func All(n int, s string) []string {
	if n > len(s) {
		return nil
	}
	retval := []string{}
	for i := 0; i <= len(s)-n; i++ {
		substr := s[i : i+n]
		retval = append(retval, substr)
	}
	return retval
}

// UnsafeFirst returns the first substring of s with length n without any
// error handling (so it's panic at the disco time).
func UnsafeFirst(n int, s string) string {
	return s[:n]
}

// First returns the first substring of s with length n, safely.
func First(n int, s string) (first string, ok bool) {
	if n > len(s) {
		first = ""
		ok = false
	} else {
		first = s[:n]
		ok = true
	}
	return first, ok
}
