package main

import (
	"fmt"
)

// https://app.codility.com/demo/results/trainingK4BM6J-KRW/
// O(N)
func Solution(A []int) int {
	if 0 == len(A) {
		return -1
	}
	value := A[0]
	size := 0
	for _, a := range A {
		if 0 == size {
			size++
			value = a
		} else {
			if a != value {
				size--
			} else {
				size++
			}
		}
	}

	candidate := -1
	if 0 != size {
		candidate = value
	}

	index := -1
	count := 0
	for i, a := range A {
		if a == candidate {
			count++
			if count > len(A)/2 {
				return i
			}
		}
	}

	return index
}

func main() {
	fmt.Println(Solution([]int{6, 8, 4, 6, 8, 6, 6}))     // 6
	fmt.Println(Solution([]int{1, 2, 19, 1, 4, 2, 1}))    // -1
	fmt.Println(Solution([]int{3, 4, 3, 2, 3, -1, 3, 3})) // 7
}
