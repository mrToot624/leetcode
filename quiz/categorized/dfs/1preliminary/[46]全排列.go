package main

//leetcode submit region begin(Prohibit modification and deletion)
func permute(nums []int) [][]int {
	var res [][]int
	var track []int
	used := make([]bool, len(nums))
	backtrack_46(nums, track, used, &res)
	return res
}

func backtrack_46(nums, track []int, used []bool, res *[][]int) {
	if len(track) == len(nums) {
		solution := make([]int, len(track))
		copy(solution, track)
		*res = append(*res, solution)
		return
	}

	for i, num := range nums {
		if used[i] {
			continue
		}

		used[i] = true
		track = append(track, num)
		backtrack_46(nums, track, used, res)
		track = track[:len(track)-1]
		used[i] = false
	}
}
//leetcode submit region end(Prohibit modification and deletion)
