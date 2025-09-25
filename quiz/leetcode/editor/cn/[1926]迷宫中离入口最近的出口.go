package main

// leetcode submit region begin(Prohibit modification and deletion)
func nearestExit(maze [][]byte, entrance []int) int {
	m, n := len(maze), len(maze[0])
	visited := make(map[[2]int]bool)
	start := [2]int(entrance)
	visited[start] = true

	q := [][2]int{start}
	var step int
	for len(q) > 0 {
		l := len(q)
		for i := 0; i < l; i++ {
			cur := q[0]
			q = q[1:]
			nextPos := getNext(cur, m, n)
			if len(nextPos) < 4 && step > 0 {
				return step
			}
			for _, pos := range nextPos {
				if maze[pos[0]][pos[1]] != '+' && !visited[pos] {
					q = append(q, pos)
					visited[pos] = true
				}
			}
		}
		step++
	}
	return -1
}

func getNext(start [2]int, m, n int) [][2]int {
	var next [][2]int
	for _, delta := range [][2]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}} {
		nextI, nextJ := start[0]+delta[0], start[1]+delta[1]
		if nextI >= 0 && nextI < m && nextJ >= 0 && nextJ < n {
			next = append(next, [2]int{nextI, nextJ})
		}
	}
	return next
}

//leetcode submit region end(Prohibit modification and deletion)
