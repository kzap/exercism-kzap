// Package proverb solves the Proverb side exercise on Exercism's Go track.
package proverb

import (
	"fmt"
)

// Proverb takes in a slice of strings and returns a slice of proverbs
func Proverb(rhyme []string) []string {
	proverbLen := len(rhyme)
	proverb := make([]string, proverbLen)

	if proverbLen == 0 {
		return proverb
	}

	if proverbLen > 1 {
		for i := 0; i < proverbLen-1; i++ {
			proverb[i] = fmt.Sprintf("For want of a %v the %v was lost.", rhyme[i], rhyme[i+1])
		}
	}

	proverb[proverbLen-1] = fmt.Sprintf("And all for the want of a %v.", rhyme[0])

	return proverb
}
