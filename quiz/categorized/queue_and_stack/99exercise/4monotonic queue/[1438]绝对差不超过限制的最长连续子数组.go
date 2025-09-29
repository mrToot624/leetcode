package main

//leetcode submit region begin(Prohibit modification and deletion)
func longestSubarray(nums []int, limit int) int {
	left, right := 0, 0
	var maxQ, minQ []int

	var res int
	for right < len(nums) {
		for len(maxQ) > 0 && maxQ[len(maxQ)-1] < nums[right] {
			maxQ = maxQ[:len(maxQ)-1]
		}
		maxQ = append(maxQ, nums[right])
		for len(minQ) > 0 && minQ[len(minQ)-1] > nums[right] {
			minQ = minQ[:len(minQ)-1]
		}
		minQ = append(minQ, nums[right])
		right++

		for len(maxQ) > 0 && len(minQ) > 0 && maxQ[0]-minQ[0] > limit && left < right {
			if maxQ[0] == nums[left] {
				maxQ = maxQ[1:]
			}
			if minQ[0] == nums[left] {
				minQ = minQ[1:]
			}
			left++
		}
		res = max(res, right-left)
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
