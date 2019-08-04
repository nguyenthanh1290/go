package main

import (
	"fmt"
)

// https://app.codility.com/demo/results/trainingE4989R-ZHY/
// Detected time complexity: O(log(N + M))
// Least common multiple
func eat(N, M int) int {
	if 1 == M {
		return N
	}
	if N == M {
		return 1
	}

	a := N
	b := M
	for b != 0 {
		a, b = b, a%b
	}

	return N / a
}

func main() {
	fmt.Println(eat(10, 4))   // 5
	fmt.Println(eat(10, 8))   // 5
	fmt.Println(eat(100, 4))  // 25
	fmt.Println(eat(100, 8))  // 25
	fmt.Println(eat(1000, 4)) // 125
	fmt.Println(eat(1000, 8)) // 125

	fmt.Println(eat(10, 3))     // 10
	fmt.Println(eat(10, 9))     // 10
	fmt.Println(eat(7, 9))      // 7
	fmt.Println(eat(2, 3))      // 2
	fmt.Println(eat(2, 19))     // 2
	fmt.Println(eat(2, 111119)) // 2

	fmt.Println(eat(2, 1))          // 2
	fmt.Println(eat(3, 1))          // 3
	fmt.Println(eat(3, 2))          // 3
	fmt.Println(eat(1000000000, 1)) // 1000000000
}
