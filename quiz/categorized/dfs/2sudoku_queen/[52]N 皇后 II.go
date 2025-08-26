package _sudoku_queen

// leetcode submit region begin(Prohibit modification and deletion)
func totalNQueens(n int) int {
	var res int

	board := make([][]bool, n)
	for i := 0; i < n; i++ {
		board[i] = make([]bool, n)
	}

	backtrack_52(board, &res, 0)
	return res
}

func backtrack_52(board [][]bool, res *int, level int) {
	n := len(board)
	if level == n {
		*res++
	}

	for col := 0; col < n; col++ {
		if isConflicted_52(board, level, col) {
			continue
		}

		board[level][col] = true
		backtrack_52(board, res, level+1)
		board[level][col] = false
	}
}

func isConflicted_52(board [][]bool, level, col int) bool {
	for row := 0; row < level; row++ {
		if board[row][col] {
			return true
		}
	}

	for i, j := level-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] {
			return true
		}
	}

	for i, j := level-1, col+1; i >= 0 && j < len(board); i, j = i-1, j+1 {
		if board[i][j] {
			return true
		}
	}

	return false
}

//leetcode submit region end(Prohibit modification and deletion)
