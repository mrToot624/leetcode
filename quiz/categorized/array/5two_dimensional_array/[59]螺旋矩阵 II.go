package main

// leetcode submit region begin(Prohibit modification and deletion)
func generateMatrix(n int) [][]int {
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}

	upper, lower := 0, n-1
	left, right := 0, n-1
	num := 1

	for num <= n*n {
		for j := left; j <= right; j++ {
			res[upper][j] = num
			num++
		}
		upper++

		for i := upper; i <= lower; i++ {
			res[i][right] = num
			num++
		}
		right--

		for j := right; j >= left; j-- {
			res[lower][j] = num
			num++
		}
		lower--

		for i := lower; i >= upper; i-- {
			res[i][left] = num
			num++
		}
		left++
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
