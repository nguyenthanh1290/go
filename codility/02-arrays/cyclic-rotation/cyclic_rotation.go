package main

import (
	"fmt"
)

func Solution(A []int, K int) []int {
	if K == 0 || len(A) == 0 || len(A) == 1 {
		return A
	}

	if K > len(A) {
		K = K % len(A)
	}

	arr := A[len(A)-K:]
	arr = append(arr, A[:len(A)-K]...)
	return arr
}

func main() {
	fmt.Println(Solution([]int{4, 6, 0, 0, -2, 10}, 15))
}

// https://app.codility.com/programmers/lessons/2-arrays/cyclic_rotation/