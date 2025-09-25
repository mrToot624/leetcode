package main

// leetcode submit region begin(Prohibit modification and deletion)
func maxSubarraySumCircular(nums []int) int {
	n := len(nums)
	preSum := make([]int, 2*n+1)
	for i := 0; i < len(preSum)-1; i++ {
		preSum[i+1] = preSum[i] + nums[i%n]
	}

	mq := []int{0}
	res := nums[0]
	left, right := 0, 1
	for right < len(preSum)-1 {
		for len(mq) > 0 && mq[len(mq)-1] > preSum[right] {
			mq = mq[:len(mq)-1]
		}
		mq = append(mq, preSum[right])
		right++

		if right-left > n {
			if mq[0] == preSum[left] {
				mq = mq[1:]
			}
			left++
		}
		res = max(res, preSum[right]-mq[0])
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
