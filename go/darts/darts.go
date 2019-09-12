// Package darts scores a dart throw
// √(𝑥𝑝−𝑥𝑐)^2+(𝑦𝑝−𝑦𝑐)^2 < 𝑟
package darts

import (
	"math"
)

// Score accepts the location of a
// thrown dart and returns the score
func Score(x, y float64) int {

	location := math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))

	if location > 10 {
		return 0
	} else if location <= 10 && location > 5 {
		return 1
	} else if location <= 5 && location > 1 {
		return 5
	}

	return 10
}
