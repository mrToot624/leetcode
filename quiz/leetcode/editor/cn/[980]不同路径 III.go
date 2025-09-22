package main

// leetcode submit region begin(Prohibit modification and deletion)
func uniquePathsIII(grid [][]int) int {
	var cnt, zeroCnt int
	m, n := len(grid), len(grid[0])
	var startI, startJ int
	for i, row := range grid {
		for j, num := range row {
			if num == 1 {
				startI, startJ = i, j
			} else if num == 0 {
				zeroCnt++
			}
		}
	}

	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i == m || j < 0 || j == n || grid[i][j] == -1 {
			return
		}
		if grid[i][j] == 2 {
			if zeroCnt == 0 {
				cnt++
			}
			return
		}

		prev := grid[i][j]
		grid[i][j] = -1
		if prev == 0 {
			zeroCnt--
		}
		dfs(i-1, j)
		dfs(i, j-1)
		dfs(i+1, j)
		dfs(i, j+1)
		if prev == 0 {
			zeroCnt++
		}
		grid[i][j] = prev
	}
	dfs(startI, startJ)
	return cnt
}

//leetcode submit region end(Prohibit modification and deletion)
