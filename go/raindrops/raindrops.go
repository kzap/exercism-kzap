// Package raindrops solves the Raindrops side exercise on Exercism's Go track.
package raindrops

import (
	"fmt"
	"strings"
)

// Convert takes a number and outputs a string, the contents of which depend on the number's factors:
//
// If the number has 3 as a factor, output 'Pling'.
// If the number has 5 as a factor, output 'Plang'.
// If the number has 7 as a factor, output 'Plong'.
//
// Else return the number
func Convert(i int) string {
	var sb strings.Builder

	if i%3 == 0 {
		sb.Grow(5)
		sb.WriteString("Pling")
	}

	if i%5 == 0 {
		sb.Grow(5)
		sb.WriteString("Plang")
	}

	if i%7 == 0 {
		sb.Grow(5)
		sb.WriteString("Plong")
	}

	if sb.Len() > 0 {
		return sb.String()
	}

	return fmt.Sprint(i)
}
