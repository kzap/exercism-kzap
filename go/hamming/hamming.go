package hamming

import (
	"fmt"
)

func Distance(a, b string) (int, error) {
	if (len(a) != len(b)) {
		return 0, fmt.Errorf("Length of strings do not match")
	}

	var count int

	for key, letterA := range a {
		if (letterA != rune(b[key])) {
			count++
		}
	}

	return count, nil
}
