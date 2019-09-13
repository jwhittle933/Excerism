package hamming

import "errors"

// Distance func compares each letter in strings a and b
// and returns the number of differences
func Distance(a, b string) (int, error) {

	if len(a) != len(b) {
		return 0, errors.New("invalid")
	}

	count := 0

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			count++
		}
	}

	return count, nil
}
