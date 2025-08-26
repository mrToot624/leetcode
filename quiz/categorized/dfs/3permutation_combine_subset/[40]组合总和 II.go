package main

import "sort"

//leetcode submit region begin(Prohibit modification and deletion)
func combinationSum2(candidates []int, target int) [][]int {
	var res [][]int
	var track []int

	sort.Ints(candidates)
	backtrack_40(track, candidates, &res, target, 0, 0)
	return res
}

func backtrack_40(track, candidates []int, res *[][]int, target, curSum, start int)  {
	if curSum >= target {
		if curSum == target {
			combination := make([]int, len(track))
			copy(combination, track)
			*res = append(*res, combination)
		}
		return
	}

	for i := start; i < len(candidates); i++ {
		if i > start && candidates[i] == candidates[i-1] {
			continue
		}

		curSum += candidates[i]
		track = append(track, candidates[i])
		backtrack_40(track, candidates, res, target, curSum, i+1)
		track = track[:len(track)-1]
		curSum -= candidates[i]
	}
}
//leetcode submit region end(Prohibit modification and deletion)
