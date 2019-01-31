package space

import "strings"

//Age takes a planet and seconds and returns how old someone would be on planet
func Age(seconds float64, planet string) float64 {
	var age float64
	switch strings.ToLower(planet) {
	case "earth":
		age = seconds / 31557600.00
	case "mercury":
		age = seconds / 0.2408467
	case "venus":
		age = seconds / 0.61519726
	case "mars":
		age = seconds / 1.8808158
	case "jupiter":
		age = seconds / 11.862615
	case "saturn":
		age = seconds / 29.447498
	case "uranus":
		age = seconds / 84.016846
	case "neptune":
		age = seconds / 164.79132
	}
	return age
}
