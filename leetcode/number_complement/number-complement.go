package bitwise

func findComplement(num int) int {
	result := 0
	power := 1

	for num > 0 {
		result += (num%2 ^ 1) * power

		power <<= 1 // power *= 2
		num >>= 1   // num /= 2
	}

	return result
}
