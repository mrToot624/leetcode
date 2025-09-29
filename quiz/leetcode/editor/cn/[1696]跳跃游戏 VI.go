package main

// leetcode submit region begin(Prohibit modification and deletion)
func maxResult(nums []int, k int) int {
	n := len(nums)
	dp := make([]int, n)
	dp[0] = nums[0]
	mq := []int{nums[0]}
	for i := 1; i < n; i++ {
		dp[i] = mq[0] + nums[i]
		for len(mq) > 0 && dp[i] > mq[len(mq)-1] {
			mq = mq[:len(mq)-1]
		}
		mq = append(mq, dp[i])
		if i-k >= 0 && mq[0] == dp[i-k] {
			mq = mq[1:]
		}
	}
	return dp[n-1]
}

//leetcode submit region end(Prohibit modification and deletion)
