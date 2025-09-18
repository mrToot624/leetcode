package main

// leetcode submit region begin(Prohibit modification and deletion)
func lengthOfLIS(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	dp[0] = 1
	for i := 1; i < n; i++ {
		maxL := 0
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				maxL = max(maxL, dp[j])
			}
		}
		dp[i] = maxL + 1
	}
	maxL := 1
	for i := range dp {
		maxL = max(maxL, dp[i])
	}
	return maxL
}

//leetcode submit region end(Prohibit modification and deletion)
