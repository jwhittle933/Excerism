package collatzconjecture

import (
	"errors"
)

// CollatzConjecture iterates 
func CollatzConjecture(input int) (int, error) {

	if input <= 0 {
		return 0, errors.New("zero or a negative number")
	}

	count := 0

	for {
		if input == 1 {
			break
		}

		if input & 1 == 1 {
			input = (3*input)+1
			count++
		}

		input = input / 2
		count++

	}

	return count, nil
}
