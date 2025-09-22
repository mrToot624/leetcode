package main

// leetcode submit region begin(Prohibit modification and deletion)
func numSubarrayProductLessThanK(nums []int, k int) int {
	left, right := 0, 0
	prod, res := 1, 0
	for right < len(nums) {
		prod *= nums[right]
		right++

		for prod >= k && left < right {
			prod /= nums[left]
			left++
		}
		res += right-left
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
