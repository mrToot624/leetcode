package main

// leetcode submit region begin(Prohibit modification and deletion)
func maxSubArray(nums []int) int {
	res, prev := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if prev < 0 {
			prev = nums[i]
		} else {
			prev += nums[i]
		}
		res = max(res, prev)
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
