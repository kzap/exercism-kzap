// Package strand solves the RNA Transcription side exercise on Exercism's Go track.
package strand

import (
	"strings"
)

var (
	dnaNucleotides = []string{"G", "C", "T", "A"}
	rnaNucleotides = []string{"C", "G", "A", "U"}
)

// ToRNA converts DNA nucleotides into RNA nucleotides
func ToRNA(dna string) string {
	var sb strings.Builder

	for _, b := range []byte(dna) {
		s := string(b)
		for i, dnaNucleotide := range dnaNucleotides {
			if s == dnaNucleotide {
				sb.WriteString(rnaNucleotides[i])
				break
			}
		}
	}

	return sb.String()
}

// ToDNA converts RNA nucleotides into RNA nucleotides
func ToDNA(rna string) string {
	var sb strings.Builder

	for _, b := range []byte(rna) {
		s := string(b)
		for i, rnaNucleotide := range rnaNucleotides {
			if s == rnaNucleotide {
				sb.WriteString(dnaNucleotides[i])
				break
			}
		}
	}

	return sb.String()
}
