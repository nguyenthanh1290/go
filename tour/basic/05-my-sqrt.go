package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := float64(1)

	for {
		oldZ := z

		z -= (z*z - x) / (2 * z)

		fmt.Println("z = ", z)

		if oldZ == z {
			return z
		}
	}

	return z
}

func main() {
	x := float64(2222)

	fmt.Println("My sqrt:", Sqrt(x))
	fmt.Println("Go sqrt:", math.Sqrt(x))
}
