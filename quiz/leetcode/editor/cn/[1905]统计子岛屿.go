package main

//leetcode submit region begin(Prohibit modification and deletion)
func countSubIslands(grid1 [][]int, grid2 [][]int) int {
    m, n := len(grid1), len(grid1[0])

	var dfs func(i, j int, grid [][]int)
	dfs = func(i, j int, grid [][]int) {
		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == 0 {
			return
		}

		grid[i][j] = 0
		dfs(i-1, j, grid)
		dfs(i, j-1, grid)
		dfs(i+1, j, grid)
		dfs(i, j+1, grid)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid2[i][j] == 1 && grid1[i][j] == 0 {
				dfs(i, j, grid2)
			}
		}
	}

	var cnt int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid2[i][j] == 1 {
				cnt++
				dfs(i, j, grid2)
			}
		}
	}
	return cnt
}
//leetcode submit region end(Prohibit modification and deletion)
