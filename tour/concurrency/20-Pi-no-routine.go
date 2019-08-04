package main

import (
	"fmt"
	"math"
)

func main() {
	pi := 0.0

	for i := 0; i <= 5000; i++ {
		pi += 4 * math.Pow(-1, float64(i)) / (2*float64(i) + 1)
	}

	fmt.Println(pi)
}
