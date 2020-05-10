package judge

func findJudge(N int, trust [][]int) int {
	count := make([]int, N+1)

	for _, t := range trust {
		count[t[0]]--
		count[t[1]]++
	}

	for i := 1; i <= N; i++ {
		if count[i] == (N - 1) {
			return i
		}
	}

	return -1
}
