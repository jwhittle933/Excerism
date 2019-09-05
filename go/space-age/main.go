package space

// Planet type
type Planet string

const earthSeconds float64 = 31557600

// Orbital Period by Earth year
var orbPeriod = map[Planet]float64{
	"Earth":   31557600,
	"Mercury": toSeconds(0.2408467),
	"Venus":   toSeconds(0.61519726),
	"Mars":    toSeconds(1.8808158),
	"Jupiter": toSeconds(11.862615),
	"Saturn":  toSeconds(29.447489),
	"Uranus":  toSeconds(84.016846),
	"Neptune": toSeconds(164.79132),
}

// Age func
// Acept random number sec (seconds) and calculate
// number of planet years
func Age(secOfLife float64, p Planet) float64 {

	// To find age, divide total seconds of life
	// by planet seconds per year
	return secOfLife / orbPeriod[p]
}

func toSeconds(relativeEarthYears float64) float64 {
	// Convert relative earth years to planet days in year
	localDaysInYear := relativeEarthYears * 365.25
	// Planet days in year * 24 yields hours
	// Hours * 60 yields minutes
	// Minutes * 60 yields seconds
	return localDaysInYear * 24 * 60 * 60

}
