// Adds 1,000,000,000 seconds to the time given and returns a new time.
package gigasecond

import "time"

// AddGigasecond takes in a time and adds 1,000,000,000 seconds to it.
// It returns a time value after a gigasecond
func AddGigasecond(t time.Time) time.Time {
	
	gigasecond := time.Duration(1000000000)

	t = t.Add(time.Second * gigasecond)

	return t
}
