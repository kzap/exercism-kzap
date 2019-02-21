// Package etl solves the ETL side exercise on Exercism's Go track.
package etl

import (
	"strings"
)

// InputMap a list of scrabble scores
type InputMap map[int][]string

// OutputMap is a map of transformed scrabble scores
type OutputMap map[string]int

// Transform takes an input and produces an output
func Transform(input InputMap) OutputMap {
	output := make(OutputMap)

	for i, a := range input {
		for _, s := range a {
			output[strings.ToLower(s)] = i
		}
	}

	return output
}
