package protein

import (
	"errors"
	"strings"
	"unicode"
)

// ErrStop is an error given if a STOP Codon is found.
var ErrStop = errors.New("STOP Codon Found")

// ErrInvalidBase is an error given if an invalid Codon is found.
var ErrInvalidBase = errors.New("invalid Codon")

// FromRNA translates a RNA sequence into a slice of proteins
func FromRNA(s string) ([]string, error) {
	var codon string
	var rna []string

	a := []rune(s)
	for i, r := range a {
		codon = codon + string(unicode.ToUpper(r))
		if i > 0 && (i+1)%3 == 0 {
			protein, err := FromCodon(codon)
			if err != nil {
				if err == ErrStop {
					return rna, nil
				}

				return rna, err
			}

			rna = append(rna, protein)

			codon = ""
		}
	}

	return rna, nil
}

// FromCodon translates Codon into its polypeptide
func FromCodon(s string) (string, error) {
	switch strings.ToUpper(s) {
	case "AUG":
		return "Methionine", nil
	case "UUU", "UUC":
		return "Phenylalanine", nil
	case "UUA", "UUG":
		return "Leucine", nil
	case "UCU", "UCC", "UCA", "UCG":
		return "Serine", nil
	case "UAU", "UAC":
		return "Tyrosine", nil
	case "UGU", "UGC":
		return "Cysteine", nil
	case "UGG":
		return "Tryptophan", nil
	case "UAA", "UAG", "UGA":
		return "", ErrStop
	}

	return "", ErrInvalidBase
}
