package main

//leetcode submit region begin(Prohibit modification and deletion)
func removeElement(nums []int, val int) int {
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}

//leetcode submit region end(Prohibit modification and deletion)
