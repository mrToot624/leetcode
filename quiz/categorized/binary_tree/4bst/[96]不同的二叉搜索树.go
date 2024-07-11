package main

// leetcode submit region begin(Prohibit modification and deletion)
func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1

	for i := 2; i <= n; i++ {
		var sum int
		for j := 1; j <= (i+1)/2; j++ {
			countWithJAsRoot := dp[j-1] * dp[i-j]
			if i%2 == 1 && j == (i+1)/2 {
				sum += countWithJAsRoot
			} else {
				sum += 2 * countWithJAsRoot
			}
		}
		dp[i] = sum
	}
	return dp[n]
}

//leetcode submit region end(Prohibit modification and deletion)
