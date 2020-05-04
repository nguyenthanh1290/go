package fibonacci

func fib(n int) int {
	// return recursion(n)

	// memo := make([]int, n+1)
	// return memoize(n, memo)

	// return bottomUpSpaceOn(n)

	return bottomUpSpaceO1(n)
}

// Time complexity: O(2^n)
// Space complexity: O(n)
func recursion(n int) int {
	var result int

	if n == 1 || n == 2 {
		result = 1

	} else {
		result = recursion(n-1) + recursion(n-2)
	}

	return result
}

// Time complexity: O(n)
// Space complexity: O(n)
func memoize(n int, memo []int) int {
	if memo[n] != 0 {
		return memo[n]
	}

	var result int

	if n == 1 || n == 2 {
		result = 1

	} else {
		result = memoize(n-1, memo) + memoize(n-2, memo)
	}

	memo[n] = result

	return result
}

// Time complexity: O(n)
// Space complexity: O(n)
func bottomUpSpaceOn(n int) int {
	if n == 1 || n == 2 {
		return 1
	}

	bup := make([]int, n+1)
	bup[1] = 1
	bup[2] = 1
	for i := 3; i <= n; i++ {
		bup[i] = bup[i-1] + bup[i-2]
	}

	return bup[n]

}

// Time complexity: O(n)
// Space complexity: O(1)
func bottomUpSpaceO1(n int) int {
	pre := 1
	cur := 1

	for i := 3; i <= n; i++ {
		new := pre + cur
		pre, cur = cur, new
	}

	return cur

}
