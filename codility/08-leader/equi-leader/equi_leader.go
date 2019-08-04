package main

import (
	"fmt"
)

// https://app.codility.com/demo/results/training78FUEM-78H/
// O(N)
func Solution(A []int) int {
	if 0 == len(A) {
		return 0
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
	} else {
		return 0
	}

	equi := 0
	left := 0
	right := 0
	for i, a := range A {
		if a == leader {
			left++
			right = count - left
		}
		if (left > (i+1)/2) && (right > (len(A)-(i+1))/2) {
			equi++
		}
	}

	return equi
}

func main() {
	fmt.Println(Solution([]int{4, 3, 4, 4, 4, 2})) // 2
}
