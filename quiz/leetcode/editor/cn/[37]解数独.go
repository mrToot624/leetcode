package main

// leetcode submit region begin(Prohibit modification and deletion)
func solveSudoku(board [][]byte) {
	var rowSet, colSet [9][9]bool
	var blockSet [3][3][9]bool
	var pos [][2]int

	for i, row := range board {
		for j, numByte := range row {
			if numByte == '.' {
				pos = append(pos, [2]int{i, j})
			} else {
				num := numByte - '1'
				rowSet[i][num] = true
				colSet[j][num] = true
				blockSet[i/3][j/3][num] = true
			}
		}
	}

	var backtrack func(int) bool
	backtrack = func(p int) bool {
		if p == len(pos) {
			return true
		}

		i, j := pos[p][0], pos[p][1]
		for k := 0; k < 9; k++ {
			if rowSet[i][k] || colSet[j][k] || blockSet[i/3][j/3][k] {
				continue
			}

			board[i][j] = byte(k + '1')
			rowSet[i][k] = true
			colSet[j][k] = true
			blockSet[i/3][j/3][k] = true

			if backtrack(p + 1) {
				return true
			}

			rowSet[i][k] = false
			colSet[j][k] = false
			blockSet[i/3][j/3][k] = false
		}
		return false
	}
	backtrack(0)
}

//leetcode submit region end(Prohibit modification and deletion)
