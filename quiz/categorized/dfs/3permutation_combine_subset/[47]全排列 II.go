package main

import "sort"

//leetcode submit region begin(Prohibit modification and deletion)
func permuteUnique(nums []int) [][]int {
	var res [][]int
	var track []int
	used := make([]bool, len(nums))

	sort.Ints(nums)
	backtrack_47(track, nums, used, &res)
	return res
}

func backtrack_47(track, nums []int, used []bool, res *[][]int) {
	if len(track) == len(nums) {
		permutation := make([]int, len(track))
		copy(permutation, track)
		*res = append(*res, permutation)
	}

	for i, num := range nums {
		if used[i] || (i > 0 && nums[i] == nums[i-1] && !used[i-1]) {
			continue
		}

		track = append(track, num)
		used[i] = true
		backtrack_47(track, nums, used, res)
		used[i] = false
		track = track[:len(track)-1]
	}
}
//leetcode submit region end(Prohibit modification and deletion)
