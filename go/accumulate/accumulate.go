package accumulate

// A binary function for folding integers.
type converterFunc func(string) string

func Accumulate(list []string, fn converterFunc) []string {
	newList := make([]string, Length(list))

	for i, v := range list {
		newList[i] = fn(v)
	}

	return newList
}

// Length takes in a slice of integers and returns number of values in the slice.
func Length(list []string) (length int) {
	for range list {
		length++
	}

	return
}
