// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym solves the Acronym exercise on Exercism's Go track.
package acronym

import (
	"strings"
	"unicode"
)

// Abbreviate takes a string and returns the first letter of each word.
func Abbreviate(s string) string {

	var abbreviation strings.Builder

	for _, word := range SplitWords(s) {
		abbreviation.WriteString(strings.ToUpper(word)[:1])
	}

	return abbreviation.String()
}

// SplitWords splits a string into words of Letters and Numbers
func SplitWords(s string) []string {
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}

	return strings.FieldsFunc(s, f)
}
