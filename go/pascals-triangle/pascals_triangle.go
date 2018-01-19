package pascal

const testVersion = 1

// Triangle calculates a Pascals Triangle of height n.
func Triangle(n int) [][]int {
	retval := [][]int{}
	for current := 0; current < n; current++ {
		curRow := make([]int, current+1)
		for i := range curRow {
			if i == 0 {
				curRow[i] = 1
			} else if i == len(curRow)-1 {
				curRow[i] = 1
			} else {
				prevRow := retval[current-1]
				curRow[i] = prevRow[i-1] + prevRow[i]
			}
		}
		retval = append(retval, curRow)
	}
	return retval
}
