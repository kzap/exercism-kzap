// Package space solves the Space Age exercise on Exercism's Go track.
package space

import (
//	"math"
)

type Planet string
type PlanetDetail struct {
	name       Planet
	earthYears float64
}

var (
	secondsPerEarthYear float64 = 31557600
	PlanetDetails               = []PlanetDetail{
		{
			name:       "Earth",
			earthYears: 1,
		},
		{
			name:       "Mercury",
			earthYears: 0.2408467,
		},
		{
			name:       "Venus",
			earthYears: 0.61519726,
		},
		{
			name:       "Mars",
			earthYears: 1.8808158,
		},
		{
			name:       "Jupiter",
			earthYears: 11.862615,
		},
		{
			name:       "Saturn",
			earthYears: 29.447498,
		},
		{
			name:       "Uranus",
			earthYears: 84.016846,
		},
		{
			name:       "Neptune",
			earthYears: 164.79132,
		},
	}
)

// Age takes in an age in seconds, calculate how old someone would be on a planet:
// - Earth: orbital period 365.25 Earth days, or 31557600 seconds
// - Mercury: orbital period 0.2408467 Earth years
// - Venus: orbital period 0.61519726 Earth years
// - Mars: orbital period 1.8808158 Earth years
// - Jupiter: orbital period 11.862615 Earth years
// - Saturn: orbital period 29.447498 Earth years
// - Uranus: orbital period 84.016846 Earth years
// - Neptune: orbital period 164.79132 Earth years
func Age(seconds float64, planet Planet) (planetYears float64) {

	planetYears = seconds / secondsPerEarthYear

	for _, planetDetail := range PlanetDetails {
		if planetDetail.name == planet {
			planetYears /= planetDetail.earthYears
			break
		}
	}

	return
}
