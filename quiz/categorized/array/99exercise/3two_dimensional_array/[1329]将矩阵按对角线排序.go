package main

import "sort"

// leetcode submit region begin(Prohibit modification and deletion)
func diagonalSort(mat [][]int) [][]int {
	m, n := len(mat), len(mat[0])
	sortDiagonal := func(i0, j0 int) {
		var temp []int
		for i, j := i0, j0; i < m && j < n; i, j = i+1, j+1 {
			temp = append(temp, mat[i][j])
		}
		sort.Ints(temp)
		for k, num := range temp {
			mat[i0+k][j0+k] = num
		}
	}

	for i := 0; i < m-1; i++ {
		sortDiagonal(i, 0)
	}
	for j := 1; j < n-1; j++ {
		sortDiagonal(0, j)
	}
	return mat
}

//leetcode submit region end(Prohibit modification and deletion)
