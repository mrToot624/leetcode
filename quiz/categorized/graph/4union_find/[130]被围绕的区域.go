package main

// leetcode submit region begin(Prohibit modification and deletion)
func solve(board [][]byte) {
	m, n := len(board), len(board[0])
	uf := NewUnionFind(m*n + 1)
	dummy := m * n
	flatten := func(i, j int) int {
		return i*n + j
	}
	for i := 0; i < m; i++ {
		if board[i][0] == 'O' {
			uf.Union(flatten(i, 0), dummy)
		}
		if board[i][n-1] == 'O' {
			uf.Union(flatten(i, n-1), dummy)
		}
	}
	for j := 0; j < n; j++ {
		if board[0][j] == 'O' {
			uf.Union(flatten(0, j), dummy)
		}
		if board[m-1][j] == 'O' {
			uf.Union(flatten(m-1, j), dummy)
		}
	}

	dict := [][]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			if board[i][j] == 'O' {
				for _, delta := range dict {
					adjI, adjJ := i+delta[0], j+delta[1]
					if board[adjI][adjJ] == 'O' {
						uf.Union(flatten(i, j), flatten(i+delta[0], j+delta[1]))
					}
				}
			}
		}
	}

	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			if board[i][j] == 'O' {
				if !uf.Connected(flatten(i, j), dummy) {
					board[i][j] = 'X'
				}
			}
		}
	}
}

//leetcode submit region end(Prohibit modification and deletion)
