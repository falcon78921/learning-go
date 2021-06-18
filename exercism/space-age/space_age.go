// space is a package that calculates age on different planetary bodies using a seconds input
package space

import "fmt"

// Function that calculates planetary age based on inputs
func Age(ageSeconds float64, planet string) float64 {

	var orbitSeconds float64

	switch orbitSeconds {
	case Mercury:
		7595341.5312
	case Venus:
		19400860.79136
	case Earth:
		31557600
	case Mars:
		59313407.0688
	case Jupiter:
		374099426.64
	case Saturn:
		928656296.928
	case Uranus:
		2649555255.46
	case Neptune:
		5196859067.52
	default:
		fmt.Println("Please provide a value for planet.")
	}

	planetaryAge := ageSeconds / orbitSeconds
	return float64(int(planetaryAge*100)) / 100
}
