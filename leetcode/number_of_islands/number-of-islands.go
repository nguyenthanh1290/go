package islands

// Time Complexity: O(N), where N is the number of pixels in the image. We might process every pixel.
// Space Complexity: O(N), the size of the implicit call stack when calling dfs.
func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	num := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				num++
				dfs(grid, i, j) // sink the entire island
			}
		}
	}

	return num
}

func dfs(grid [][]byte, i, j int) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[i]) || grid[i][j] == '0' {
		return
	}

	grid[i][j] = '0'

	dfs(grid, i+1, j)
	dfs(grid, i-1, j)
	dfs(grid, i, j+1)
	dfs(grid, i, j-1)

	return
}
