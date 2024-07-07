package main

// leetcode submit region begin(Prohibit modification and deletion)
func removeDuplicates(nums []int) int {
	slow, fast := 0, 1
	for fast < len(nums) {
		if nums[slow] != nums[fast] {
			slow++
			nums[slow] = nums[fast]
		}
		fast++
	}
	return slow + 1
}

//leetcode submit region end(Prohibit modification and deletion)
