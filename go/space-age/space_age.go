package space

// Planet name
type Planet string

const earth Planet = "Earth"
const earthSecondsYear = 31557600

var orbitalPeriod = map[Planet]float64{
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	earth:     1,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

// Age calcs a person age given the seconds of life and a planet
func Age(seconds float64, planet Planet) float64 {
	op := orbitalPeriod[planet]
	return seconds / (earthSecondsYear * op)
}
