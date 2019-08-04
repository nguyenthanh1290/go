package main

import (
	"fmt"
	"sort"
)

func main() {
	a1 := []int{2, 1, 3, 5, 4}
	a2 := []int{2, 3, 4, 1, 5}
	a3 := []int{1, 3, 4, 2, 5}
	a4 := []int{1, 2, 6, 3, 4, 5}
	fmt.Println(Solution(a1))
	fmt.Println(Solution(a2))
	fmt.Println(Solution(a3))
	fmt.Println(Solution(a4))
}

func Solution(A []int) int {
	count := 0

	for i := 0; i < len(A); i++ {
		arr := A[:i+1]
		if areAllBulbsTurnedOn(arr) {
			count++
		}
	}

	return count
}

// given the input condition: the elements of A are all distinct,
// we can verify if all the bulbs are turned of by sorting and
// then check if the first element is 1 and
// the last element equals to the total number of the array.
func areAllBulbsTurnedOn(arr []int) bool {
	sort.Ints(arr)
	if arr[0] == 1 && arr[len(arr)-1] == len(arr) {
		return true
	}
	return false
}
