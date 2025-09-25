package main

//leetcode submit region begin(Prohibit modification and deletion)
func longestOnes(nums []int, k int) int {
    left, right := 0, 0
	var zeroCnt int
	var res int
	for right < len(nums) {
		if nums[right] == 0 {
			zeroCnt++
		}
		right++
		for zeroCnt > k && left < right {
			if nums[left] == 0 {
				zeroCnt--
			}
			left++
		}
		res = max(res, right-left)
	}
	return res
}
//leetcode submit region end(Prohibit modification and deletion)
