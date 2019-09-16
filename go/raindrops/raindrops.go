package raindrops

import "strconv"

// Convert accepts input integer,
// checks if input has 3, 5, or 7 as factor,
// and returns a string
func Convert(input int) string {
	var s string

	if input%3 == 0 {
		s += "Pling"
	}

	if input%5 == 0 {
		s += "Plang"
	}

	if input%7 == 0 {
		s += "Plong"
	}

	if s == "" {
		return strconv.Itoa(input)
	}

	return s
}
