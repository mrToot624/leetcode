package main

//leetcode submit region begin(Prohibit modification and deletion)
func combinationSum3(k int, n int) [][]int {
	var res [][]int
	var track []int

	backtrack_216(track, &res, 0, n, k, 1)
	return res
}

func backtrack_216(track []int, res *[][]int, curSum, target, maxNum, start int)  {
	if curSum >= target || len(track) >= maxNum {
		if curSum == target && len(track) == maxNum {
			combination := make([]int, len(track))
			copy(combination, track)
			*res = append(*res, combination)
		}
		return
	}

	for i := start; i <= 9; i++ {
		track = append(track, i)
		curSum += i
		backtrack_216(track, res, curSum, target, maxNum, i+1)
		curSum -= i
		track = track[:len(track)-1]
	}
}
//leetcode submit region end(Prohibit modification and deletion)
