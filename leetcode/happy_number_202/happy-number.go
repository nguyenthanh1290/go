package happy_number_202

func isHappy(n int) bool {
	seen := make(map[int]struct{})

	for {
		sum := sumOfSquareDigits(n)

		if sum == 1 {
			return true

		} else if _, ok := seen[sum]; ok {
			return false

		} else {
			seen[sum] = struct{}{}
		}

		n = sum
	}
}

func sumOfSquareDigits(n int) int {
	sum := 0

	for n != 0 {
		i := n % 10
		sum += i * i

		n /= 10
	}

	return sum
}
