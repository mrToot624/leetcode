package main

// leetcode submit region begin(Prohibit modification and deletion)
func removeDuplicates(nums []int) int {
	slow, fast := 0, 0
	for fast < len(nums) {
		var count int
		if fast+1 < len(nums) && nums[fast] == nums[fast+1] {
			v := nums[fast]
			for fast < len(nums) && nums[fast] == v {
				count++
				if count <= 2{
					nums[slow] = nums[fast]
					slow++
				}
				fast++
			}
		} else {
			nums[slow] = nums[fast]
			slow++
			fast++
		}
	}
	return slow
}

//leetcode submit region end(Prohibit modification and deletion)
