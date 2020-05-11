package judge

func findJudge(N int, trust [][]int) int {
	// return findJudgeGraph(N, trust)
	return findJudgeSimplify(N, trust)
}

func findJudgeGraph(N int, trust [][]int) int {
	ins := make([]int, N+1)
	outs := make([]int, N+1)

	for _, edge := range trust {
		outs[edge[0]]++
		ins[edge[1]]++
	}

	for i := 1; i <= N; i++ {
		if ins[i] == (N-1) && outs[i] == 0 {
			return i
		}
	}

	return -1
}

func findJudgeSimplify(N int, trust [][]int) int {
	count := make([]int, N+1)

	for _, t := range trust {
		count[t[0]] = -1

		if count[t[1]] != -1 {
			count[t[1]]++
		}
	}

	for i := 1; i <= N; i++ {
		if count[i] == (N - 1) {
			return i
		}
	}

	return -1
}
