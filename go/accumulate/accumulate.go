package accumulate

const testVersion = 1

// Accumulate applies the operation to the collection of strings, returning
// the result
func Accumulate(input []string, operation func(string) string) []string {
	return accumulateAppend(input, operation)
	// return accumulateArray(input, operation)
}

// accumulateAppend applies the operation to the collection of strings
// using append to build the result set.
// ok  	_/Users/ferlatte/exercism/go/accumulate	1.193s
func accumulateAppend(input []string, operation func(string) string) []string {
	var res []string
	for _, item := range input {
		res = append(res, operation(item))
	}
	return res
}

// ok  	_/Users/ferlatte/exercism/go/accumulate	2.513s
func accumulateArray(input []string, operation func(string) string) []string {
	res := make([]string, len(input))
	for i, item := range input {
		res[i] = operation(item)
	}
	return res
}
