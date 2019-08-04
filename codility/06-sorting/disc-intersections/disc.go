package disc

const limit = 10000000

// O(N) complexity and O(N) memory
// https://app.codility.com/demo/results/training39GYB5-J3P/
// Detected time complexity:
// O(N*log(N)) or O(N)
func Solution(A []int) int {
	totalDiscs := len(A)
	startPts := make([]int, totalDiscs)
	endPts := make([]int, totalDiscs)
	for i, a := range A {
		start := i - a
		if start < 0 {
			start = 0
		}

		end := i + a
		if end > (totalDiscs - 1) {
			end = totalDiscs - 1
		}

		startPts[start]++
		endPts[end]++
	}

	result := 0
	activeDiscs := 0
	for i := 0; i < totalDiscs; i++ {
		if startPts[i] > 0 {
			result += activeDiscs * startPts[i]
			result += startPts[i] * (startPts[i] - 1) / 2

			if result > limit {
				return -1
			}

			activeDiscs += startPts[i]
		}
		activeDiscs -= endPts[i]
	}

	return result
}
