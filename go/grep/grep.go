// Package grep solves the Grep side exercise on Exercism's Go track.
package grep

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

// Flags are grep modifiers
type Flags map[string]bool

// Search does takes in a grep pattern and flags and searches through files
// and returns a slice of string results.
func Search(pattern string, flags, filenames []string) []string {
	results := []string{}

	// parse flags
	flagsActivated := ParseFlags(flags)
	showLineNumber := IsFlagActive("n", flagsActivated)
	showFileNameOnly := IsFlagActive("l", flagsActivated)
	caseInsensitive := IsFlagActive("i", flagsActivated)
	inverseMatch := IsFlagActive("v", flagsActivated)
	matchExact := IsFlagActive("x", flagsActivated)

	//make regexp pattern
	var sb strings.Builder
	if caseInsensitive {
		sb.WriteString("(?i)")
	}
	if matchExact {
		sb.WriteString("^")
	}
	sb.WriteString(pattern)
	if matchExact {
		sb.WriteString("$")
	}
	regExpPattern := sb.String()

	for _, filename := range filenames {
		var lineNum int

		f, err := os.Open(filename)
		if err != nil {
			log.Fatalf("error opening file: %v - %v\n", filename, err)
			os.Exit(1)
		}
		defer f.Close()

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			line := scanner.Text()
			lineFound := false
			lineNum++

			matched, err := regexp.MatchString(regExpPattern, line)
			if err != nil {
				log.Fatalf("regexp error using [%v]: %v\n", pattern, err)
				os.Exit(1)
			}
			if matched != inverseMatch {
				lineFound = true
			}

			if lineFound {
				var sb strings.Builder

				if showFileNameOnly {
					if StringInSlice(filename, results) {
						continue
					}

					sb.Grow(len(filename))
					sb.WriteString(filename)
				} else {
					// if we are searching for more than file, prepend filename
					if len(filenames) > 1 {
						sb.Grow(len(filename) + 1)
						sb.WriteString(filename)
						sb.WriteString(":")
					}

					if showLineNumber {
						lineNumStr := fmt.Sprintf("%d", lineNum)
						sb.Grow(len(lineNumStr) + 1)
						sb.WriteString(lineNumStr)
						sb.WriteString(":")
					}

					sb.Grow(len(line))
					sb.WriteString(line)
				}

				line = sb.String()
				results = append(results, line)
			}
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}

	return results
}

// ParseFlags returns a map of flags used
func ParseFlags(flags []string) Flags {
	flagsActivated := make(Flags)

	for _, s := range flags {
		s = strings.TrimLeft(s, "-")
		flagsActivated[s] = true
	}

	return flagsActivated
}

// IsFlagActive takes in a flag and a map of Flags and returns a bool
func IsFlagActive(flag string, flagsActivated Flags) bool {
	_, ok := flagsActivated[flag]
	return ok
}

// StringInSlice checks if a string exists in a list of strings.
func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}

	return false
}
