package minavg

// Detected time complexity:
// O(N ** 2)
func Solution1(A []int) int {
	result := 0
	minAvg := 100000.0

	for i := 0; i < len(A)-1; i++ {
		sum := float64(A[i])
		divisor := 2.0
		for j := i + 1; j < len(A); j++ {
			sum += float64(A[j])
			avg := sum / divisor

			if avg < minAvg {
				minAvg = avg
				result = i
			}

			divisor++
		}
	}

	return result
}

// https://app.codility.com/demo/results/training75V4F8-ETS/
// Detected time complexity:
// O(N)
// The slice contains at least two elements =>
// we only need to check for slices of length 2 and 3,
// because slices of length 4 are actually 2 slices of length 2.
func Solution(A []int) int {
	minAvg := float64(A[0]+A[1]) / 2.0
	result := 0

	for i := 0; i < len(A)-1; i++ {
		avg := float64(A[i]+A[i+1]) / 2.0
		if avg < minAvg {
			minAvg = avg
			result = i
		}
	}

	for i := 0; i < len(A)-2; i++ {
		avg := float64(A[i]+A[i+1]+A[i+2]) / 3.0
		if avg < minAvg {
			minAvg = avg
			result = i
		}
	}

	return result
}
