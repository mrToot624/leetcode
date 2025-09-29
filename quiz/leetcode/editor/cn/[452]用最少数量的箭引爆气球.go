package main

import "sort"

//leetcode submit region begin(Prohibit modification and deletion)
func findMinArrowShots(points [][]int) int {
    sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})
	cnt := 1
	right := points[0][1]
	for i := 1; i < len(points); i++ {
		if points[i][0] > right {
			right = points[i][1]
			cnt++
		}
	}
	return cnt
}
//leetcode submit region end(Prohibit modification and deletion)
