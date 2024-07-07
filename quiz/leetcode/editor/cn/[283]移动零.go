package main

// leetcode submit region begin(Prohibit modification and deletion)
func moveZeroes(nums []int) {
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast] != 0 {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}

	for i := slow; i < len(nums); i++ {
		nums[i] = 0
	}
}

//leetcode submit region end(Prohibit modification and deletion)
