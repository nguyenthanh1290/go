package main

import (
	"fmt"
)

// https://app.codility.com/demo/results/training8E34M3-HNE/
// Detected time complexity: O(sqrt(N))
func count(N int) int {
	c := 0
	i := 1
	for i*i < N {
		if N%i == 0 {
			c += 2
		}
		i++
	}
	if i*i == N {
		c++
	}

	return c
}

func main() {
	fmt.Println(count(24)) // 8
}
