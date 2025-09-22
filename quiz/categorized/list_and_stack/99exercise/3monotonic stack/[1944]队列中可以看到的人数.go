package main

// leetcode submit region begin(Prohibit modification and deletion)
func canSeePersonsCount(heights []int) []int {
	n := len(heights)
	res := make([]int, n)
	var ms []int
	for i := n - 1; i >= 0; i-- {
		var pCnt int
		for len(ms) > 0 && heights[i] > ms[len(ms)-1] {
			pCnt++
			ms = ms[:len(ms)-1]
		}
		if len(ms) > 0 {
			pCnt++
		}
		res[i] = pCnt
		ms = append(ms, heights[i])
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
