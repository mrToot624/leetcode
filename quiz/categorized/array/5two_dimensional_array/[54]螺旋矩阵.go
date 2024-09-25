package main

// leetcode submit region begin(Prohibit modification and deletion)
func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	upper, lower := 0, m
	left, right := 0, n

	var res []int
	for len(res) < m*n {
		for j := left; j <= right; j++ {
			res = append(res, matrix[upper][j])
		}
		upper++

		for i := upper; i <= lower; i++ {
			res = append(res, matrix[i][right])
		}
		right--

		for j := right; j >= left; j-- {
			res = append(res, matrix[lower][j])
		}
		lower--

		for i := lower; i >= upper; i++ {
			res = append(res, matrix[i][left])
		}
		left++
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
