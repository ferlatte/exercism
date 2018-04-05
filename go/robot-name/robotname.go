package robotname

import (
	"fmt"
	"math/rand"
)

// Robot is, well, a robot. With a name.
type Robot struct {
	name string
}

// Name returns the name of the robot, generating it if necessary
func (r *Robot) Name() string {
	if r.name == "" {
		r.name = makeName()
	}
	return r.name
}

// Reset factory resets the robot.
func (r *Robot) Reset() {
	r.name = ""
}

var names map[string]bool

func makeName() string {
	var try func() string
	try = func() string {
		letters := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
		i := rand.Intn(len(letters))
		prefix := string(letters[i])
		i = rand.Intn(len(letters))
		prefix = prefix + string(letters[i])
		i = rand.Intn(999)
		name := prefix + fmt.Sprintf("%03d", i)
		return name
	}
	if names == nil {
		names = make(map[string]bool)
	}
	for {
		n := try()
		if !names[n] {
			names[n] = true
			return n
		}
	}
}
