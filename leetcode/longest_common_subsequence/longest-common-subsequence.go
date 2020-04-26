package subsequence

func longestCommonSubsequence(text1 string, text2 string) int {
	// return recursion(text1, text2, len(text1), len(text2))

	// m := len(text1)
	// n := len(text2)
	// memo := make([][]int, m+1)
	// for i := 0; i < m+1; i++ {
	// 	memo[i] = make([]int, n+1)
	// 	for j := 0; j < n+1; j++ {
	// 		memo[i][j] = -1
	// 	}
	// }
	// return memoize(text1, text2, m, n, memo)

	return bottomUp(text1, text2)
}

// Time complexity: O(2^(N+M))
// Space complexity: O(1)
func recursion(p, q string, m, n int) int {
	var result int

	if m == 0 || n == 0 {
		result = 0

	} else if p[m-1] == q[n-1] {
		result = 1 + recursion(p, q, m-1, n-1)

	} else {
		c1 := recursion(p, q, m-1, n)
		c2 := recursion(p, q, m, n-1)
		result = max(c1, c2)
	}

	return result
}

func memoize(p, q string, m, n int, memo [][]int) int {
	if memo[m][n] != -1 {
		return memo[m][n]
	}

	var result int

	if m == 0 || n == 0 {
		result = 0

	} else if p[m-1] == q[n-1] {
		result = 1 + memoize(p, q, m-1, n-1, memo)

	} else {
		c1 := memoize(p, q, m-1, n, memo)
		c2 := memoize(p, q, m, n-1, memo)
		result = max(c1, c2)
	}

	memo[m][n] = result

	return result
}

func bottomUp(p, q string) int {
	m := len(p)
	n := len(q)
	memo := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		memo[i] = make([]int, n+1)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if p[i] == q[j] {
				memo[i+1][j+1] = 1 + memo[i][j]

			} else {
				memo[i+1][j+1] = max(memo[i+1][j], memo[i][j+1])
			}
		}
	}

	return memo[m][n]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
