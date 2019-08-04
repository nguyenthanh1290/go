package main

import (
	"fmt"
)

func Solution(A []int) int {
	m := make(map[int]int)
	for _, v := range A {
		m[v]++
	}
	for k := range m {
		if m[k]%2 != 0 {
			return k
		}
	}
	return -1
}

func main() {
	fmt.Println(Solution([]int{9, 3, 9, 3, 9, 7, 9}))
}

// https://app.codility.com/programmers/lessons/2-arrays/odd_occurrences_in_array/