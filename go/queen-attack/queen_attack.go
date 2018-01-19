package queenattack

import (
	"errors"
	"strconv"
	"strings"
	"unicode/utf8"
)

const testVersion = 2

type position struct {
	x int
	y int
}

// CanQueenAttack takes the position of two Queens as strings using
// http://en.wikipedia.org/wiki/Algebraic_notation_(chess) and returns
// true if so.
func CanQueenAttack(w, b string) (ok bool, err error) {
	if w == b {
		return false, errors.New("positions must be different")
	}
	wq, err := notationToPosition(w)
	if err != nil {
		return false, err
	}
	bq, err := notationToPosition(b)
	if err != nil {
		return false, err
	}
	if wq.x == bq.x || wq.y == bq.y {
		return true, nil
	}
	slope := float64(wq.y-bq.y) / float64(wq.x-bq.x)
	if slope == 1.0 || slope == -1.0 {
		return true, nil
	}

	return false, nil
}

func notationToPosition(n string) (p position, err error) {
	retval := position{}

	if len(n) != 2 {
		return retval, errors.New("Bad notation string")
	}
	xy := strings.Split(n, "")
	x, err := xToInt(xy[0])
	if err != nil {
		return retval, err
	}
	y, err := yToInt(xy[1])
	if err != nil {
		return retval, err
	}
	retval.x = x
	retval.y = y
	return retval, nil
}

func yToInt(y string) (int, error) {
	retval, err := strconv.Atoi(y)
	if err != nil {
		return retval, errors.New("Bad y position")
	}
	if retval < 1 || retval > 8 {
		return retval, errors.New("Y position is not 1-8")
	}
	return retval, nil
}

func xToInt(x string) (int, error) {
	retval := 0
	r, _ := utf8.DecodeRuneInString(x)
	if r == utf8.RuneError {
		return retval, errors.New("Rune error in notation string")
	}
	// x must be a-h
	if !('a' <= r && r <= 'h') {
		return 0, errors.New("X position is not a-h")
	}
	retval = int(r-'a') + 1
	return retval, nil
}
