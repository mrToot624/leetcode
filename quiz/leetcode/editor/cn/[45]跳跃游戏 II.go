package main

// leetcode submit region begin(Prohibit modification and deletion)
func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	if nums[0] >= len(nums)-1 {
		return 1
	}
	var cur, cnt int
	for cur < len(nums)-1 {
		nextCur, farther := cur, cur+nums[cur]
		for i := cur + 1; i <= cur+nums[cur]; i++ {
			if farther < i+nums[i] {
				farther = i + nums[i]
				nextCur = i
			}
			if farther >= len(nums)-1 {
				return cnt + 2
			}
		}
		cur = nextCur
		cnt++
	}
	return cnt
}

//leetcode submit region end(Prohibit modification and deletion)
