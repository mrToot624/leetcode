package main

// leetcode submit region begin(Prohibit modification and deletion)
func shortestPathBinaryMatrix(grid [][]int) int {
	n := len(grid)
	if grid[0][0] == 1 || grid[n-1][n-1] == 1 {
		return -1
	}
	if n == 1 {
		return 1
	}
	q := [][2]int{{0, 0}}
	grid[0][0] = 1
	directions := [][2]int{
		{-1, 0}, {0, -1}, {1, 0}, {0, 1},
		{-1, -1}, {1, -1}, {1, 1}, {-1, 1},
	}
	var cnt int
	for len(q) > 0 {
		l := len(q)
		for i := 0; i < l; i++ {
			cur := q[0]
			q = q[1:]
			x, y := cur[0], cur[1]
			for _, direction := range directions {
				x1, y1 := x+direction[0], y+direction[1]
				if x1 == n-1 && y1 == n-1 {
					return cnt + 2
				}
				if x1 >= 0 && x1 < n && y1 >= 0 && y1 < n && grid[x1][y1] == 0 {
					q = append(q, [2]int{x1, y1})
					grid[x1][y1] = 1
				}
			}
		}
		cnt++
	}
	return -1
}

//leetcode submit region end(Prohibit modification and deletion)
