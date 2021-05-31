// space is a package that calculates age on different planetary bodies using a seconds input
package main

import "fmt"

// Function that calculates planetary age based on inputs
func Age(ageSeconds float64, planet string) float64 {

    var orbitSeconds float64

    if planet == "Mercury" {
        orbitSeconds = 7595341.5312
    } else if planet == "Venus" {
        orbitSeconds = 19400860.79136
    } else if planet == "Earth" {
        orbitSeconds = 31557600
    } else if planet == "Mars" {
        orbitSeconds = 59313407.0688
    } else if planet == "Jupiter" {
        orbitSeconds = 374099426.64
    } else if planet == "Saturn" {
        orbitSeconds = 928656296.928
    } else if planet == "Uranus" {
        orbitSeconds = 2649555255.46
    } else if planet == "Neptune" {
        orbitSeconds = 5196859067.52
    }

    planetaryAge := ageSeconds / orbitSeconds
    return float64(int(planetaryAge * 100)) / 100
}

func main() {
    // Test Age() with Earth
    earthAge := Age(1000000000, "Earth")
    fmt.Println(earthAge)
}
