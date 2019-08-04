package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %f", float64(e))
}

const e = 1e-8

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	if x == 0 {
		return 0, nil
	}

	z := float64(1)

	for {
		oldZ := z

		z -= (z*z - x) / (2 * z)

		if math.Abs(z-oldZ) < e {
			return z, nil
		}
	}

	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
	fmt.Println(Sqrt(-2))
}
