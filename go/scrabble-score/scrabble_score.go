// Package scrabble solves the Scrabble exercise on Exercism's Go track.
package scrabble

import (
	"strings"
)

// LetterScore is a mapping from a Letter to its scrabble score.
type LetterScore map[rune]int

// Score takes in a word and returns its Scrabble score.
func Score(word string) int {
	letterScore := LetterScore{
		'A': 1,
		'E': 1,
		'I': 1,
		'O': 1,
		'U': 1,
		'L': 1,
		'N': 1,
		'R': 1,
		'S': 1,
		'T': 1,
		'D': 2,
		'G': 2,
		'B': 3,
		'C': 3,
		'M': 3,
		'P': 3,
		'F': 4,
		'H': 4,
		'V': 4,
		'W': 4,
		'Y': 4,
		'K': 5,
		'J': 8,
		'X': 8,
		'Q': 10,
		'Z': 10,
	}

	if len(word) == 0 {
		return 0
	}

	var score int
	for _, r := range strings.ToUpper(word) {
		if _, exists := letterScore[r]; exists {
			score += letterScore[r]
		}
	}

	return score
}
