// Package gigasecond should have a package comment that summarizes what it's about.
package gigasecond

// import path for the time package from the standard library
import "time"

const giga time.Duration = 1000000000

// AddGigasecond returns the input time
// plus a gigasecond
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * giga)
}
