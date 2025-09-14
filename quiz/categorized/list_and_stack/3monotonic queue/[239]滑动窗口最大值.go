package main

// leetcode submit region begin(Prohibit modification and deletion)
func maxSlidingWindow(nums []int, k int) []int {
	res := make([]int, len(nums)-k+1)
	var mq []int
	for i, num := range nums {
		for len(mq) > 0 && mq[len(mq)-1] < nums[i] {
			mq = mq[:len(mq)-1]
		}
		mq = append(mq, num)
		if i >= k-1 {
			res[i-k+1] = mq[0]
			if nums[i-k+1] == mq[0] {
				mq = mq[1:]
			}
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
