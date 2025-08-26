package main

// leetcode submit region begin(Prohibit modification and deletion)
func closedIsland(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	var cnt int

	var dfs func(row, col int, grid [][]int)
	dfs = func(i, j int, grid [][]int) {
		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == 1 {
			return
		}

		grid[i][j] = 1

		dfs(i-1, j, grid)
		dfs(i, j-1, grid)
		dfs(i+1, j, grid)
		dfs(i, j+1, grid)
	}

	for j := 0; j < n; j++ {
		if grid[0][j] == 0 {
			dfs(0, j, grid)
		}
		if grid[m-1][j] == 0 {
			dfs(m-1, j, grid)
		}
	}
	for i := 0; i < m; i++ {
		if grid[i][0] == 0 {
			dfs(i, 0, grid)
		}
		if grid[i][n-1] == 0 {
			dfs(i, n-1, grid)
		}
	}

	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			if grid[i][j] == 0 {
				cnt++
				dfs(i, j, grid)
			}
		}
	}
	return cnt
}

//leetcode submit region end(Prohibit modification and deletion)
