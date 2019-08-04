package main

import (
	"fmt"
	"math"
	"sort"
)

func maxRepeated(A, B int) int {
	lower := int(math.Ceil(math.Sqrt(float64(A))))
	upper := int(math.Floor(math.Sqrt(float64(B))))

	sqrts := []int{}
	for i := lower; i <= upper; i++ {
		sqrts = append(sqrts, primitiveSqrt(i))
	}

	sort.Ints(sqrts)

	return sqrts[len(sqrts)-1]
}

func primitiveSqrt(n int) int {
	steps := 0
	p := int(math.Sqrt(float64(n)))
	if p*p == n {
		steps++
		return steps + primitiveSqrt(p)
	}
	return steps
}

func main() {
	fmt.Println(maxRepeated(10, 20))
	fmt.Println(maxRepeated(6000, 7000))
}
