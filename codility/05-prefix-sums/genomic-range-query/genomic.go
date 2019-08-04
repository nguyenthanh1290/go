package genomic

import (
	"strings"
)

// O(N + M)
// https://app.codility.com/demo/results/trainingJ65WK3-RK2/
// Detected time complexity: O(N + M)
func Solution1(S string, P []int, Q []int) []int {
	var genoms [3][]int
	genoms[0] = make([]int, len(S)+1) // A
	genoms[1] = make([]int, len(S)+1) // C
	genoms[2] = make([]int, len(S)+1) // G

	for i, v := range S {
		a, c, g := 0, 0, 0
		switch v {
		case 'A':
			a = 1
		case 'C':
			c = 1
		case 'G':
			g = 1
		}
		genoms[0][i+1] = genoms[0][i] + a
		genoms[1][i+1] = genoms[1][i] + c
		genoms[2][i+1] = genoms[2][i] + g
	}

	result := make([]int, len(P))
	for i := 0; i < len(P); i++ {
		fromIndex := P[i]
		toIndex := Q[i] + 1

		if genoms[0][toIndex]-genoms[0][fromIndex] > 0 {
			result[i] = 1
		} else if genoms[1][toIndex]-genoms[1][fromIndex] > 0 {
			result[i] = 2
		} else if genoms[2][toIndex]-genoms[2][fromIndex] > 0 {
			result[i] = 3
		} else {
			result[i] = 4
		}
	}

	return result
}

// Failed: extreme_large (all max ranges)
// https://app.codility.com/demo/results/training73WDJH-RUG/
// Detected time complexity: O(N + M)
func Solution(S string, P []int, Q []int) []int {
	result := make([]int, len(P))
	for i := 0; i < len(P); i++ {
		strip := S[P[i] : Q[i]+1]
		if strings.Contains(strip, "A") {
			result[i] = 1
		} else if strings.Contains(strip, "C") {
			result[i] = 2
		} else if strings.Contains(strip, "G") {
			result[i] = 3
		} else {
			result[i] = 4
		}
	}

	return result
}
