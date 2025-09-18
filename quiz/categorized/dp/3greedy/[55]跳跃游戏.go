package main

// leetcode submit region begin(Prohibit modification and deletion)
func canJump(nums []int) bool {
	right := 0
	for i, num := range nums {
		if i <= right {
			right = max(right, i+num)
			if right >= len(nums)-1 {
				return true
			}
		} else {
			return false
		}
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)
