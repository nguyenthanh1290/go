package square

func isPerfectSquare(num int) bool {
	// return multiply(num)
	// return binarySearch(num)
	return babylonian(num)
	// return newton(num)
}

// Time complexity : O(sqrt(N))
// Space complexity : O(1)
func multiply(num int) bool {
	for i := 1; i*i <= num; i++ {
		if i*i == num {
			return true
		}
	}
	return false
}

// Time complexity : O(logN)
// Space complexity : O(1)
func binarySearch(num int) bool {
	left := 1
	right := num

	for left <= right {
		mid := left + (right-left)/2

		square := mid * mid
		if square == num {
			return true

		} else if square > num {
			right = mid - 1

		} else {
			left = mid + 1
		}
	}

	return false
}

// https://en.wikipedia.org/wiki/Methods_of_computing_square_roots#Babylonian_method
func babylonian(num int) bool {
	digits := 0
	temp := num
	for temp != 0 {
		digits++
		temp /= 10
	}
	// digits = (digits / 2) + 1

	// guest := float32(1)
	// for ; digits >= 0; digits-- {
	// 	guest *= 10
	// }

	guest := float32(digits * 100)

	var prev int
	for int(guest) != prev {
		prev = int(guest)
		guest = 0.5 * (guest + float32(num)/guest)
	}

	if prev*prev == num {
		return true
	}
	if (prev+1)*(prev+1) == num {
		return true
	}

	return false
}

// https://en.wikipedia.org/wiki/Newton%27s_method#Square_root_of_a_number
func newton(num int) bool {
	if num == 1 {
		return true
	}

	guest := num / 2

	for guest*guest > num {
		guest = (guest + num/guest) / 2
	}

	return guest*guest == num
}
