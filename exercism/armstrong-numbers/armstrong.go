package armstrong

import (
	"math"
	"strconv"
)

// IsNumber checks whether a number is an Armstrong number,
// that is the sum of its own digits each raised to the power of the number of digits.
func IsNumber(n int) bool {
	return isNumberVb(n)
}

func isNumberVa(n int) bool {
	sum := 0
	s := strconv.Itoa(n)
	y := float64(len(s))

	for _, v := range s {
		x, _ := strconv.Atoi(string(v))
		sum += int(math.Pow(float64(x), y))
	}
	
	if sum == n {
		return true
	}
	return false
}

func isNumberVb(n int) bool {
	y := float64(int(math.Log10(float64(n))) + 1)

	sum := 0
	i := n
	for {
		x := float64(i % 10)
		sum += int(math.Pow(x, y))
		i /= 10
		if i == 0 {
			break
		}
	}

	if sum == n {
		return true
	}
	return false
}