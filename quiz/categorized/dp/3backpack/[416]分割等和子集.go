package main

// leetcode submit region begin(Prohibit modification and deletion)
func canPartition(nums []int) bool {
	sum, n := 0, len(nums)
	for _, num := range nums {
		sum += num
	}
	if sum%2 == 1 || n == 1 {
		return false
	}
	w := sum / 2

	dp := make([]bool, w+1)
	dp[0] = true
	for i := 1; i <= n; i++ {
		for j := w; j >= nums[i-1]; j-- {
			dp[j] = dp[j] || dp[j-nums[i-1]]
		}
	}
	return dp[w]
}

//leetcode submit region end(Prohibit modification and deletion)
