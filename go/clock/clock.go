package clock

import "time"

const testVersion = 4

// You can find more details and hints in the test file.

// Clock keeps track of time.
type Clock time.Time

// New makes a Clock
func New(hour, minute int) Clock {
	// If you make a new Clock(45,20), the
	// time library will correctly wrap into the next day. This makes
	// things look correct, but equality tests fail because the Time
	// isn't actually the same. So we make a Time from our inputs,
	// extract the correct hour, minute values from it, and then
	// make a new Time from the sanitized hour, minute.
	t := time.Date(0, time.January, 0, hour, minute, 0, 0, time.UTC)
	hour, minute, _ = t.Clock()
	t = time.Date(0, time.January, 0, hour, minute, 0, 0, time.UTC)
	c := Clock(t)
	return c
}

func (c Clock) String() string {
	t := time.Time(c)
	return t.Format("15:04")
}

// Add adds minutes to the clock. You can add negative minutes.
func (c Clock) Add(minutes int) Clock {
	d := time.Duration(minutes) * time.Minute
	t := time.Time(c)
	t = t.Add(d)
	return Clock(t)
}
