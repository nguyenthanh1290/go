package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	// func
	result := sum(233, 23)
	fmt.Println("result = ", result)

	//
	result2, err := sqrt(-14)
	if err != nil {
		fmt.Println("err: ", err)
	} else {
		fmt.Println("square = ", result2)
	}
}

// When two or more consecutive named function parameters share a type,
// you can omit the type from all but the last.
func sum(x, y int) int {
	return x + y
}

// A function can return any number of results.
func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Undefied for nagative numbers")
	}

	return math.Sqrt(x), nil
}
