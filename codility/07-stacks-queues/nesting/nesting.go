package main

import (
	"fmt"
)

// Performance issue:
// https://app.codility.com/demo/results/trainingD69EKZ-5BP/
// Detected time complexity: O(N)

// https://app.codility.com/demo/results/trainingZEWMVA-SUN/
// Detected time complexity: O(N)
func Solution(S string) int {
	if 0 == len(S) {
		return 1
	}
	if 0 != len(S)%2 {
		return 0
	}

	stack := []rune{}
	for _, r := range S {
		switch r {
		case '(':
			stack = append(stack, r)
		case ')':
			len := len(stack)
			if 0 == len {
				return 0
			}
			if stack[len-1] != '(' {
				return 0
			}
			stack = stack[:len-1]
		default:
			return 0
		}
	}

	if 0 == len(stack) {
		return 1
	}
	return 0
}

func main() {
	fmt.Println(Solution(""))           // 1
	fmt.Println(Solution("("))          // 0
	fmt.Println(Solution("(()(())())")) // 1
	fmt.Println(Solution("())"))        // 0
}
