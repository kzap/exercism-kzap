// Package isogram solves the Isogram exercise on Exercism's Go track.
package isogram

import (
	"unicode"
)

// IsIsogram takes in a string and returns true if it is an Isogram or false if not.
func IsIsogram(s string) bool {
	m := make(map[rune]bool)

	for _, r := range s {
		if !unicode.IsLetter(r) {
			continue
		}

		upperRune := unicode.ToUpper(r)

		if m[upperRune] {
			return false
		}

		m[upperRune] = true
	}

	return true
}
