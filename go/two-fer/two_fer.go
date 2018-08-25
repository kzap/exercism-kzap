package twofer

import "fmt"

// ShareWith should have a comment documenting it.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}

	response := fmt.Sprintf("One for %s, one for me.", name)

	return response
}
