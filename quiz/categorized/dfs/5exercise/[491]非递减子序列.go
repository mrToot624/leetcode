package main

//leetcode submit region begin(Prohibit modification and deletion)
func findSubsequences(nums []int) [][]int {
	var res [][]int
	var track []int

	var dfs func(start int)
	dfs = func(start int) {
		if len(track) >= 2 {
			subSeq := make([]int, len(track))
			copy(subSeq, track)
			res = append(res, subSeq)
		}

		used := make(map[int]bool)
		for i := start; i < len(nums); i++ {
			if (len(track) > 0 && nums[i] < track[len(track)-1]) || used[nums[i]] {
				continue
			}
			used[nums[i]] = true
			track = append(track, nums[i])
			dfs(i+1)
			track = track[:len(track)-1]
		}
	}
	dfs(0)
	return res
}
//leetcode submit region end(Prohibit modification and deletion)
