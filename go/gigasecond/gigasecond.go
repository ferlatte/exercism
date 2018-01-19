package gigasecond

// import path for the time package from the standard library
import "time"

const testVersion = 4

// AddGigasecond adds 10^9 seconds to t and returns the new time.Time
func AddGigasecond(t time.Time) time.Time {
	gigasecond := time.Duration(1000000000) * time.Second
	return t.Add(gigasecond)
}
