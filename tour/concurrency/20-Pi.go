// http://www.mathscareers.org.uk/article/calculating-pi/
// Calculating Pi (π) using infinite series
// Gregory-Leibniz Series
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(pi(5000))
}

func pi(n int) float64 {
	c := make(chan float64)
	for i := 0; i <= n; i++ {
		go term(c, float64(i))
	}

	pi := 0.0
	for i := 0; i <= n; i++ {
		pi += <-c
	}

	return pi
}

func term(c chan float64, k float64) {
	c <- 4 * math.Pow(-1, k) / (2*k + 1)
}
