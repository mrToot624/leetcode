package main

//leetcode submit region begin(Prohibit modification and deletion)
func numEnclaves(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	var cnt int
	var needCnt bool

	var dfs func(row, col int, grid [][]int)
	dfs = func(i, j int, grid [][]int) {
		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == 0 {
			return
		}

		if needCnt {
			cnt++
		}
		grid[i][j] = 0

		dfs(i-1, j, grid)
		dfs(i, j-1, grid)
		dfs(i+1, j, grid)
		dfs(i, j+1, grid)
	}

	for j := 0; j < n; j++ {
		if grid[0][j] == 1 {
			dfs(0, j, grid)
		}
		if grid[m-1][j] == 1 {
			dfs(m-1, j, grid)
		}
	}
	for i := 0; i < m; i++ {
		if grid[i][0] == 1 {
			dfs(i, 0, grid)
		}
		if grid[i][n-1] == 1 {
			dfs(i, n-1, grid)
		}
	}

	needCnt = true
	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			if grid[i][j] == 1 {
				dfs(i, j, grid)
			}
		}
	}
	return cnt
}
//leetcode submit region end(Prohibit modification and deletion)
