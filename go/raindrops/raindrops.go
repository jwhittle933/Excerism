package raindrops

import "strconv"


// Convert ...
func Convert(input int) string {
	var s string
	f := map[int]bool{1: true}

	for i := 2; i <= input; i++ {
		if input%i == 0 {
			f[i] = true
		}
	}

	if !f[3] && !f[5] && !f[7] {
		return strconv.Itoa(input)
	}

	if f[3] {
		s = "Pling"
	}

	if f[5] {
		s += "Plang"
	}

	if f[7] {
		s += "Plong"
	}

	return s
}
