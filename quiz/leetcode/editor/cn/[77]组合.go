package main

//leetcode submit region begin(Prohibit modification and deletion)
func combine(n int, k int) [][]int {
	var res [][]int
	var track []int
	backtrack_77(track, &res, n, k, 1)
	return res
}

func backtrack_77(track []int, res *[][]int, n, k, start int) {
	if len(track) == k {
		combination := make([]int, k)
		copy(combination, track)
		*res = append(*res, combination)
	}

	for i := start; i <= n; i++ {
		track = append(track, i)
		backtrack_77(track, res, n, k, i+1)
		track = track[:len(track)-1]
	}
}
//leetcode submit region end(Prohibit modification and deletion)
