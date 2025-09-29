package main

import "sort"

// leetcode submit region begin(Prohibit modification and deletion)
func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	right := intervals[0]
	var cnt int
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < right[1] {
			cnt++
		} else {
			right = intervals[i]
		}
	}
	return cnt
}

//leetcode submit region end(Prohibit modification and deletion)
