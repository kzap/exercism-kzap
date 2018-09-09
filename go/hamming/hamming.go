// Package hamming solves the Hamming exercise on Exercism's Go track.
package hamming

import (
	"errors"
)

// Distance takes in 2 strings and returns the Hamming distance between them.
// It is found by comparing two DNA strands and counting how many of the
// nucleotides are different from their equivalent in the other string.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("length of strings do not match")
	}

	var count int
	for i, s := range []byte(a) {
		if s != b[i] {
			count++
		}
	}

	return count, nil
}
