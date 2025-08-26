package _sudoku_queen

import "strings"

// leetcode submit region begin(Prohibit modification and deletion)
func solveNQueens(n int) [][]string {
	var res [][]string

	board := make([][]string, n)
	for i := 0; i < n; i++ {
		row := make([]string, n)
		for j := 0; j < n; j++ {
			row[j] = "."
		}
		board[i] = row
	}

	backtrack_51(board, &res, 0)
	return res
}

func backtrack_51(board [][]string, res *[][]string, level int) {
	n := len(board)
	if level == n {
		var compacted []string
		for _, row := range board {
			compacted = append(compacted, strings.Join(row, ""))
		}
		*res = append(*res, compacted)
		return
	}

	for col := 0; col < n; col++ {
		if isConflicted_51(board, level, col) {
			continue
		}

		board[level][col] = "Q"
		backtrack_51(board, res, level+1)
		board[level][col] = "."
	}
}

func isConflicted_51(board [][]string, level, col int) bool {
	for i := 0; i < level; i++ {
		if board[i][col] == "Q" {
			return true
		}
	}

	for i, j := level-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == "Q" {
			return true
		}
	}

	for i, j := level-1, col+1; i >= 0 && j < len(board); i, j = i-1, j+1 {
		if board[i][j] == "Q" {
			return true
		}
	}

	return false
}

//leetcode submit region end(Prohibit modification and deletion)
