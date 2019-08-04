package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow)) // x^y
}

/*
- Functions are VALUES too. They can be passed around just like other values.
- Function values may be used as function ARGUMENTS and RETURN values.
*/

// 13
// 5
// 81

/*
func compute(fn func(float64, float64) float64, x, y float64) float64 {
	return fn(x, y)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot, 3, 4))
	fmt.Println(compute(math.Pow, 3, 4))
}
*/