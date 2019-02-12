// Package scrabble solves the Scrabble exercise on Exercism's Go track.
package scrabble

import (
	"unicode"
)

var m = map[rune]int{
	'A': 1, 'E': 1, 'I': 1, 'O': 1, 'U': 1, 'L': 1, 'N': 1, 'R': 1, 'S': 1, 'T': 1,
	'D': 2, 'G': 2,
	'B': 3, 'C': 3, 'M': 3, 'P': 3,
	'F': 4, 'H': 4, 'V': 4, 'W': 4, 'Y': 4,
	'K': 5,
	'J': 8, 'X': 8,
	'Q': 10, 'Z': 10,
}

// Score takes in a word and returns its Scrabble score using a map.
func Score(word string) int {

	if len(word) == 0 {
		return 0
	}

	var score int
	for _, r := range word {
		if letterScore, exists := m[unicode.ToUpper(r)]; exists {
			score += letterScore
		}
	}

	return score
}

// ScoreSwitch takes in a word and returns its Scrabble score using a switch case.
func ScoreSwitch(word string) int {

	if len(word) == 0 {
		return 0
	}

	var score int
	for _, r := range word {
		switch unicode.ToUpper(r) {
		case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
			score++
		case 'D', 'G':
			score += 2
		case 'B', 'C', 'M', 'P':
			score += 3
		case 'F', 'H', 'V', 'W', 'Y':
			score += 4
		case 'K':
			score += 5
		case 'J', 'X':
			score += 8
		case 'Q', 'Z':
			score += 10
		}
	}

	return score
}
