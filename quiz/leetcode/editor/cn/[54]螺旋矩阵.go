package main

// leetcode submit region begin(Prohibit modification and deletion)
func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	upper, lower := 0, m-1
	left, right := 0, n-1

	var res []int
	for len(res) < m*n {
		if upper <= lower {
			for j := left; j <= right; j++ {
				res = append(res, matrix[upper][j])
			}
			upper++
		}

		if left <= right {
			for i := upper; i <= lower; i++ {
				res = append(res, matrix[i][right])
			}
			right--
		}

		if upper <= lower {
			for j := right; j >= left; j-- {
				res = append(res, matrix[lower][j])
			}
			lower--
		}

		if left <= right {
			for i := lower; i >= upper; i-- {
				res = append(res, matrix[i][left])
			}
			left++
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
