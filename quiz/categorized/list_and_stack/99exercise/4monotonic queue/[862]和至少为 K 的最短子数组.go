package main

// leetcode submit region begin(Prohibit modification and deletion)
func shortestSubarray(nums []int, k int) int {
	n := len(nums)
	preSum := make([]int, n+1)
	for i := 0; i < n; i++ {
		preSum[i+1] = preSum[i] + nums[i]
	}

	left, right := 0, 1
	mq := []int{0}
	res := n + 1
	for right < len(preSum) {
		for len(mq) > 0 && preSum[right] < mq[len(mq)-1] {
			mq = mq[:len(mq)-1]
		}
		mq = append(mq, preSum[right])
		for len(mq) > 0 && preSum[right]-mq[0] >= k && left < right {
			if preSum[left] == mq[0] {
				res = min(res, right-left)
				mq = mq[1:]
			}
			left++
		}
		right++
	}
	if res == n+1 {
		return -1
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
