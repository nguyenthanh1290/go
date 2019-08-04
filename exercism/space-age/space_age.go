// Package space provides function to calculate Earth-years old.
package space

// Solar system planets
const (
	Earth   = "Earth"
	Mercury = "Mercury"
	Venus   = "Venus"
	Mars    = "Mars"
	Jupiter = "Jupiter"
	Saturn  = "Saturn"
	Uranus  = "Uranus"
	Neptune = "Neptune"
)

// Planet orbital periods
const (
	EarthObPeriod   float64 = 31557600
	MercuryObPeriod float64 = EarthObPeriod * 0.2408467
	VenusObPeriod   float64 = EarthObPeriod * 0.61519726
	MarsObPeriod    float64 = EarthObPeriod * 1.8808158
	JupiterObPeriod float64 = EarthObPeriod * 11.862615
	SaturnObPeriod  float64 = EarthObPeriod * 29.447498
	UranusObPeriod  float64 = EarthObPeriod * 84.016846
	NeptuneObPeriod float64 = EarthObPeriod * 164.79132
)

// Age calculates and returns Earth-years old.
func Age(seconds float64, planet string) float64 {
	switch planet {
	case Earth:
		return seconds / EarthObPeriod
	case Mercury:
		return seconds / MercuryObPeriod
	case Venus:
		return seconds / VenusObPeriod
	case Mars:
		return seconds / MarsObPeriod
	case Jupiter:
		return seconds / JupiterObPeriod
	case Saturn:
		return seconds / SaturnObPeriod
	case Uranus:
		return seconds / UranusObPeriod
	case Neptune:
		return seconds / NeptuneObPeriod
	}
	return 0.00
}
