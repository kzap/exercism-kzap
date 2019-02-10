package accumulate

// A function which takes a list of strings and outputs a list of strings.
type converterFunc func(string) string

// Accumulate takes in a list of strings and a function and maps over them.
func Accumulate(list []string, fn converterFunc) []string {
	newList := make([]string, Length(list))

	for i, v := range list {
		newList[i] = fn(v)
	}

	return newList
}

// Length takes in a slice of strings and returns number of values in the slice.
func Length(list []string) (length int) {
	for range list {
		length++
	}

	return
}
