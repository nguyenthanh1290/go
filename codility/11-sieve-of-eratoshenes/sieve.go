package main

import (
	"fmt"
)

func primes(n int) []int {
	marker := make([]bool, n+1)
	for p := 2; p*p <= n; p++ {
		if !marker[p] {
			for i := p * p; i <= n; i += p {
				marker[i] = true
			}
		}
	}

	prime := []int{}
	for i := 2; i <= n; i++ {
		if !marker[i] {
			prime = append(prime, i)
		}
	}
	return prime
}

func factors(n int64) []int64 {
	r := make([]int64, 0)
	for i := int64(2); i <= n; i++ {
		for n%i == 0 {
			r = append(r, i)
			n /= i
		}
	}

	return r
}

func main() {
	fmt.Println(primes(10)) // 2, 3, 5, 7
	fmt.Println(primes(17)) /// 2, 3, 5, 7, 11, 13, 17
	fmt.Println(primes(30)) // 2, 3, 5, 7, 11, 13, 17, 19, 23, 29
}
