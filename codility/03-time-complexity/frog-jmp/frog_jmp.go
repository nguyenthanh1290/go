package main

import (
	"fmt"
)

func Solution(X int, Y int, D int) int {
	dis := Y - X
	steps := dis / D
	if dis%D == 0 {
		return steps
	}
	return steps + 1
}

func main() {
	fmt.Println(Solution(10, 85, 30))
}

// https://app.codility.com/programmers/lessons/3-time_complexity/frog_jmp/
// https://app.codility.com/demo/results/trainingMJ2TT6-F38/
// Detected time complexity:
// O(1)