package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println(isPalindrome(-101))
	fmt.Println(smallest(2))
	fmt.Println(largest(2))
}

func smallest(n int) (minProduct int, factor1 int, factor2 int) {
	lowerBound := int(math.Pow10(n - 1))
	upperBound := int(math.Pow10(n) - 1)

	minProduct = math.MaxInt32

	for f1 := lowerBound; f1 <= upperBound; f1++ {
		for f2 := lowerBound; f2 <= f1; f2++ {
			p := f1 * f2
			if p > minProduct {
				break
			}
			if p < minProduct && isPalindrome(p) {
				minProduct = p
				factor1 = f1
				factor2 = f2
			}
		}
	}

	return
}

func largest(n int) (maxProduct int, factor1 int, factor2 int) {
	lowerBound := int(math.Pow10(n - 1))
	upperBound := int(math.Pow10(n) - 1)

	for f1 := upperBound; f1 >= lowerBound; f1-- {
		for f2 := upperBound; f2 >= f1; f2-- {
			p := f1 * f2
			if p < maxProduct {
				break
			}
			if p > maxProduct && isPalindrome(p) {
				maxProduct = p
				factor1 = f1
				factor2 = f2
			}
		}
	}

	return
}

func isPalindrome(n int) bool {
	return isPalindromeVc(n)
}
func isPalindromeVc(n int) bool {
	reverse := 0
	copyOfN := n
	for copyOfN != 0 {
		reverse = reverse*10 + copyOfN%10
		copyOfN /= 10

	}
	return reverse == n
}
func isPalindromeVa(n int) bool {
	if n < 0 {
		n = -n
	}
	divisor := 1
	for n/divisor >= 10 {
		divisor *= 10
	}

	for n != 0 {
		leading := n / divisor
		trailing := n % 10

		if leading != trailing {
			return false
		}

		// Removing the leading and trailing digit from number.
		n = (n % divisor) / 10

		// Reducing divisor by a factor of 2 as 2 digits are dropped.
		divisor /= 100
	}

	return true
}
func isPalindromeVb(n int) bool {
	if n < 0 {
		n = -n
	}

	s := strconv.Itoa(n)
	l := len(s)
	for i := 0; i < l/2; i++ {
		if s[i] != s[l-1-i] {
			return false
		}
	}
	return true
}
