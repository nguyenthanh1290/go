package passingcars

const max = 1000000000

// https://app.codility.com/demo/results/trainingVWFGAN-ZNG/
// Detected time complexity:
// O(N)
func PassingCars1(A []int) int {
	total := 0
	count := 0

	for i := 0; i < len(A); i++ {
		if A[i] == 0 {
			count++
		} else {
			total += count
		}

		if total > max {
			return -1
		}
	}

	return total
}

// O(n)
// https://app.codility.com/demo/results/training8HTB7M-TV9/
// Detected time complexity: O(N)
func PassingCars(A []int) int {
	total := 0
	for _, v := range A {
		total += v
	}

	passingCars := 0
	leftTotal := 0
	for _, v := range A {
		if v == 0 {
			passingCars += total - leftTotal
			if passingCars > max {
				return -1
			}
		}
		leftTotal += v
	}

	return passingCars
}
