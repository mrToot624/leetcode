package main

//leetcode submit region begin(Prohibit modification and deletion)
func subsets(nums []int) [][]int {
	var res [][]int
	var track []int
	backtrack_78(track, nums, &res, 0)
	return res
}

func backtrack_78(track, nums []int, res *[][]int, start int) {
	subset := make([]int, len(track))
	copy(subset, track)
	*res = append(*res, subset)

	for i := start; i < len(nums); i++ {
		track = append(track, nums[i])
		backtrack_78(track, nums, res, i+1)
		track = track[:len(track)-1]
	}
}
//leetcode submit region end(Prohibit modification and deletion)
