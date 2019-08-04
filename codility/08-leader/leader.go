package main

import (
	"fmt"
)

// O(n)
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

	count := 0
	for _, a := range A {
		if a == candidate {
			count++
		}
	}

	leader := -1
	if count > len(A)/2 {
		leader = candidate
	}

	return leader
}

func main() {
	fmt.Println(Solution([]int{6, 8, 4, 6, 8, 6, 6}))  // 6
	fmt.Println(Solution([]int{1, 2, 19, 1, 4, 2, 1})) // -1
}
