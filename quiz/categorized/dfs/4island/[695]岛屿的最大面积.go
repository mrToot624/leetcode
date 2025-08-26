package main

// leetcode submit region begin(Prohibit modification and deletion)
func maxAreaOfIsland(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	var maxArea int

	var dfs func(i, j int, grid [][]int) int
	dfs = func(i, j int, grid [][]int) int {
		if i < 0 || j < 0 || i >= m || j >= n || grid[i][j] == 0 {
			return 0
		}

		grid[i][j] = 0
		return dfs(i-1, j, grid) +
			dfs(i, j-1, grid) +
			dfs(i+1, j, grid) +
			dfs(i, j+1, grid) + 1
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				maxArea = max(maxArea, dfs(i, j, grid))
			}
		}
	}
	return maxArea
}

//leetcode submit region end(Prohibit modification and deletion)
