package main

import (
	"fmt"
)

func Solution(N int, A []int) []int {
	counter := make([]int, N)
	max := 0
	lastMax := 0
	for _, v := range A {
		if v >= 1 && v <= N {
			cv := counter[v-1]
			if cv < lastMax {
				cv = lastMax
			}
			cv++
			if cv > max {
				max = cv
			}
			counter[v-1] = cv
		} else if v == N+1 {
			lastMax = max
		}
	}

	for i, v := range counter {
		if v < lastMax {
			counter[i] = lastMax
		}
	}
	return counter
}

func main() {
	fmt.Println(Solution(5, []int{3, 4, 4, 6, 1, 4, 4})) // expected: [3, 2, 2, 4, 2]
}

// https://app.codility.com/programmers/lessons/4-counting_elements/max_counters/
// https://app.codility.com/demo/results/trainingAWNA36-FCR/
// Detected time complexity:
// O(N + M)
