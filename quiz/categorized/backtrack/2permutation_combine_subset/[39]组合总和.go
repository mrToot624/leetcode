package main

//leetcode submit region begin(Prohibit modification and deletion)
func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	var track []int

	backtrack_39(track, candidates, &res, target, 0, 0)
	return res
}

func backtrack_39(track, candidates []int, res *[][]int, target, curSum, start int) {
	if curSum >= target {
		if curSum == target {
			combination := make([]int, len(track))
			copy(combination, track)
			*res = append(*res, combination)
		}
		return
	}

	for i := start; i < len(candidates); i++ {
		track = append(track, candidates[i])
		curSum += candidates[i]
		backtrack_39(track, candidates, res, target, curSum, i)
		curSum -= candidates[i]
		track = track[:len(track)-1]
	}
}
//leetcode submit region end(Prohibit modification and deletion)
