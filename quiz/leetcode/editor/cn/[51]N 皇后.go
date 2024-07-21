package main

// leetcode submit region begin(Prohibit modification and deletion)
func solveNQueens(n int) [][]string {
	var res [][]string
	queens := make([]int, n)
	column, diagonal1, diagonal2 := make(map[int]bool), make(map[int]bool), make(map[int]bool)
	backtrack_51(queens, column, diagonal1, diagonal2, &res, 0)
	return res
}

func backtrack_51(queens []int, column, diagonal1, diagonal2 map[int]bool, res *[][]string, level int) {
	n := len(queens)
	if level == n {
		*res = append(*res, generateBoard(queens))
		return
	}

	for col := 0; col < n; col++ {
		if isConflicted_51(column, diagonal1, diagonal2, level, col) {
			continue
		}

		column[col] = true
		diagonal1[level-col] = true
		diagonal2[level+col] = true
		queens[level] = col
		backtrack_51(queens, column, diagonal1, diagonal2, res, level+1)
		queens[level] = -1
		delete(diagonal2, level+col)
		delete(diagonal1, level-col)
		delete(column, col)
	}
}

func isConflicted_51(column, diagonal1, diagonal2 map[int]bool, row, col int) bool {
	return column[col] || diagonal1[row-col] || diagonal2[row+col]
}

func generateBoard(queens []int) []string {
	n := len(queens)
	res := make([]string, n)
	for row, col := range queens {
		line := make([]byte, n)
		for i := 0; i < n; i++ {
			if i == col {
				line[i] = 'Q'
			} else {
				line[i] = '.'
			}
		}
		res[row] = string(line)
	}
	return res
}
//leetcode submit region end(Prohibit modification and deletion)
