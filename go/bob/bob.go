// Package bob solves the Bob exercise on the Exercism Go Track
package bob

import (
	"strings"
)

// Hey takes an string and returns a string response depending on the string input given.
// Bob answers 'Sure.' if you ask him a question.
// He answers 'Whoa, chill out!' if you yell at him.
// He answers 'Calm down, I know what I'm doing!' if you yell a question at him.
// He says 'Fine. Be that way!' if you address him without actually saying anything.
// He answers 'Whatever.' to anything else.
func Hey(remark string) string {

	remark = strings.TrimSpace(remark)

	if IsEmpty(remark) {
		return "Fine. Be that way!"
	}

	if IsQuestion(remark) && IsYelling(remark) {
		return "Calm down, I know what I'm doing!"
	}

	if IsQuestion(remark) {
		return "Sure."
	}

	if IsYelling(remark) {
		return "Whoa, chill out!"
	}

	return "Whatever."
}

// IsQuestion checks if the string ends with a question mark.
func IsQuestion(remark string) bool {
	if string(remark[len(remark)-1:]) == "?" {
		return true
	}

	return false
}

// IsYelling checks if this string is in upper case.
func IsYelling(remark string) bool {
	if remark == strings.ToUpper(remark) && remark != strings.ToLower(remark) {
		return true
	}

	return false
}

// IsEmpty checks if the string has any characters.
func IsEmpty(remark string) bool {
	if len(remark) == 0 {
		return true
	}

	return false
}
