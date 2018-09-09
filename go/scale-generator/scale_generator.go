// Package scale solves the Scale Generator exercise on Exercism's Go track.
package scale

import (
	"strings"
)

var chromaticScaleSharp = []string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}
var chromaticScaleFlat = []string{"A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab"}

var noSharpsOrFlats = []string{"C", "a"}
var useSharps = []string{"G", "D", "A", "E", "B", "F#", "e", "b", "f#", "c#", "g#", "d#"}
var useFlats = []string{"F", "Bb", "Eb", "Ab", "Db", "Gb", "d", "g", "c", "f", "bb", "eb"}

// Scale takes a tonic, or starting note, and a set of intervals,
// and generates the musical scale starting with the tonic and following the
// specified interval pattern.
//
// Scales in Western music are based on the chromatic (12-note) scale. This
// scale can be expressed as the following group of pitches:
//
// A, A#, B, C, C#, D, D#, E, F, F#, G, G#
func Scale(tonic string, interval string) []string {

	if interval == "" {
		interval = strings.Repeat("m", 12)
	}

	var chromaticScale []string

	if StringInSlice(tonic, useFlats) {
		chromaticScale = chromaticScaleFlat
	} else {
		chromaticScale = chromaticScaleSharp
	}

	// find start of scale
	tonicUpper := strings.ToUpper(tonic)
	for i, s := range chromaticScale {
		if strings.ToUpper(s) == tonicUpper {
			// rearrange scale
			chromaticScale = append(chromaticScale[i:], chromaticScale[:i]...)
			break
		}
	}

	var scale []string
	var key int = 0
	for _, r := range interval {
		scale = append(scale, chromaticScale[key])

		switch string(r) {
		case "m":
			key += 1
		case "M":
			key += 2
		case "A":
			key += 3
		}
	}

	return scale
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
