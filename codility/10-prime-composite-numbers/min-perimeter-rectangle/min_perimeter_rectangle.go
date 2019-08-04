package main

import (
	"fmt"
)

// https://app.codility.com/demo/results/trainingZHTSZ9-M3D/
// Detected time complexity: O(sqrt(N))
func min(N int) int {
	a := 1
	b := N / a
	minPeri := 2 * (a + b)

	for a*a <= N {
		if N%a == 0 {
			b = N / a
			cur := 2 * (a + b)
			if cur < minPeri {
				minPeri = cur
			}
		}
		a++
	}

	return minPeri
}

func main() {
	fmt.Println(min(30)) // 22
}
