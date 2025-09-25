package main

// leetcode submit region begin(Prohibit modification and deletion)
func sortColors(nums []int) {
	p0, p1 := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			nums[p0], nums[i] = nums[i], nums[p0]
			p0++
			if p1 < p0 {
				p1 = p0
			}
		}
		if nums[i] == 1 {
			nums[p1], nums[i] = nums[i], nums[p1]
			p1++
		}
	}
}

//leetcode submit region end(Prohibit modification and deletion)
