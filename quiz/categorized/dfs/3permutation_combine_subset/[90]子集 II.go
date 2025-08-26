package main

import "sort"

//leetcode submit region begin(Prohibit modification and deletion)
func subsetsWithDup(nums []int) [][]int {
	var res [][]int
	var track []int

	sort.Ints(nums)
	backtrack_90(track, nums, &res, 0)
	return res
}

func backtrack_90(track, nums []int, res *[][]int, start int) {
	subset := make([]int, len(track))
	copy(subset, track)
	*res = append(*res, subset)

	for i := start; i < len(nums); i++ {
		if i > start && nums[i] == nums[i-1] {
			continue
		}

		track = append(track, nums[i])
		backtrack_90(track, nums, res, i+1)
		track = track[:len(track)-1]
	}
}
//leetcode submit region end(Prohibit modification and deletion)
