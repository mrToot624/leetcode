package main

// leetcode submit region begin(Prohibit modification and deletion)
func numIslands(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	var cnt int

	var dfs func(i, j int, grid [][]byte)
	dfs = func(i, j int, grid [][]byte) {
		if i < 0 || j < 0 || i == m || j == n || grid[i][j] == '0' {
			return
		}

		grid[i][j] = '0'

		dfs(i-1, j, grid)
		dfs(i, j-1, grid)
		dfs(i+1, j, grid)
		dfs(i, j+1, grid)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				cnt++
				dfs(i, j, grid)
			}
		}
	}
	return cnt
}

//leetcode submit region end(Prohibit modification and deletion)
