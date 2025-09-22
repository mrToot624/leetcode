package main

// leetcode submit region begin(Prohibit modification and deletion)
func minFallingPathSum(matrix [][]int) int {
	n := len(matrix)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	for j := 0; j < n; j++ {
		dp[0][j] = matrix[0][j]
	}
	for i := 1; i < n; i++ {
		for j := 0; j < n; j++ {
			minPathSum := dp[i-1][j]
			if j >= 1 {
				minPathSum = min(minPathSum, dp[i-1][j-1])
			}
			if j < n-1 {
				minPathSum = min(minPathSum, dp[i-1][j+1])
			}
			dp[i][j] = minPathSum + matrix[i][j]
		}
	}
	res := dp[n-1][0]
	for j := 1; j < n; j++ {
		res = min(res, dp[n-1][j])
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
