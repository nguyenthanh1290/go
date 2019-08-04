package div

// 0 <= A ≤ B <= 2,000,000,000
// 1 <= K <= 2,000,000,000
func Solution(A int, B int, K int) int {
	if A == 0 {
		return B/K + 1 // +1 because 0 mod any number = 0
	}
	return B/K - (A-1)/K
}

// https://app.codility.com/demo/results/trainingAZHYCH-BKG/
// Detected time complexity:
// O(1)
