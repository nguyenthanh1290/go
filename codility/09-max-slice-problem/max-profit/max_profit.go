package main

import (
	"fmt"
)

// O(n)
// https://app.codility.com/demo/results/trainingUZ5J8F-4S6/
func max(A []int) int {
	minDailyPrice := 200000
	maxProfit := 0
	for _, a := range A {
		if a < minDailyPrice {
			minDailyPrice = a
		}
		if a-minDailyPrice > maxProfit {
			maxProfit = a - minDailyPrice
		}
	}
	return maxProfit
}

func main() {
	fmt.Println(max([]int{23171, 21011, 21123, 21366, 21013, 21367})) // 356
}
