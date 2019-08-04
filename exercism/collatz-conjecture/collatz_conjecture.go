package collatzconjecture

import "errors"

// CollatzConjecture takes any [n] positive number and returns the number of steps required to reach 1 by repeating divide [n] by 2, if [n] is even, else multiply [n] by 3 and add 1.
func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("positive integer is required")
	}

	steps := 0
	for n != 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = n*3 + 1
		}
		steps++
	}
	return steps, nil
}
