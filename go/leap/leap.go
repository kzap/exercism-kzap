// Package leap solves the Leap side exercise on Exercism's Go track.
package leap

// IsLeapYear checks if the given year is a leap year and returns true if so.
func IsLeapYear(year int) bool {
	if year%4 == 0 {
		if year%100 == 0 {
			if year%400 == 0 {
				return true
			}

			return false
		}

		return true
	}

	return false
}
