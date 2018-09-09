// Package collatzconjecture solves the Collatz Conjecture exercise on Exercism's Go track.
package collatzconjecture

import (
	"errors"
)

// CollatzConjecture takes any positive integer n. If n is even, divide n by 2 to get n / 2.
// If n is odd, multiply n by 3 and add 1 to get 3n + 1.
// Repeat the process indefinitely.
//
// The conjecture states that no matter which number you start with, you will always reach 1 eventually.
//
// Given a number n, return the number of steps required to reach 1.
func CollatzConjecture(input int) (int, error) {

	steps := 0

	if input < 1 {
		return 0, errors.New("input must be greater than or equal to 1")
	}

	for {
		if input <= 1 {
			break
		}

		if input%2 == 0 {
			input = input / 2
		} else {
			input = (3 * input) + 1
		}

		steps++
	}

	return steps, nil
}
