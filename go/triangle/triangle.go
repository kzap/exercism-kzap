// Package space solves the Triange core exercise on Exercism's Go track.
package triangle

import (
	"math"
)

// Returns a kind of triagle
type Kind string

const (
	NaT = "NaT" // not a triangle
	Equ = "Equ" // equilateral
	Iso = "Iso" // isosceles
	Sca = "Sca" // scalene
)

// KindFromSides takes the lengths of a triange and returns what kind of triangle it is.
func KindFromSides(a, b, c float64) Kind {
	var k Kind

	if IsNaT(a, b, c) {
		return NaT
	}

	if IsEqu(a, b, c) {
		return Equ
	}

	if IsIso(a, b, c) {
		return Iso
	}

	if IsSca(a, b, c) {
		return Sca
	}

	return k
}

// IsNaT checks if the lengths given make a valid triangle
func IsNaT(a, b, c float64) bool {

	if a <= 0 || b <= 0 || c <= 0 {
		return true
	}

	if math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) {
		return true
	}

	if math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0) {
		return true
	}

	if !(a+b >= c) {
		return true
	}

	if !(a+c >= b) {
		return true
	}

	if !(b+c >= a) {
		return true
	}

	return false
}

// IsEqu checks if all sides are equal
func IsEqu(a, b, c float64) bool {
	if a == b && b == c {
		return true
	}

	return false
}

// IsIso checks if all 2 sides are equal
func IsIso(a, b, c float64) bool {
	if a == b || b == c || a == c {
		return true
	}

	return false
}

// IsSca checks if all 3 sides are not equal
func IsSca(a, b, c float64) bool {
	if a != b && b != c && a != c {
		return true
	}

	return false
}
