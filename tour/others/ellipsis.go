package main

import (
	"fmt"
)

func main() {
	// 1. variadic parameter
	fmt.Println(Sum(2, 3, 6))

	// 2.
	primes := []int{2, 3, 5, 7}
	fmt.Println(Sum(primes...)) // unpacking

	// 3. Array literals
	// the ... notation specifies a length equal to the number of elements in the literal
	stooges := [...]string{"Moe", "Larry", "Curly"}
	fmt.Println(len(stooges))
}

// Variadic function parameters
func Sum(nums ...int) int {
	res := 0
	for _, n := range nums {
		res += n
	}
	return res
}

// Three dots