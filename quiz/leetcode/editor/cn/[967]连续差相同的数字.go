package main

import "math"

// leetcode submit region begin(Prohibit modification and deletion)
func numsSameConsecDiff(n int, k int) []int {
	var res []int

	var dfs func(track []int)
	dfs = func(track []int) {
		if len(track) == n {
			res = append(res, list2Num(track))
			return
		}

		lastDigit := track[len(track)-1]

		if lastDigit+k <= 9 {
			track = append(track, lastDigit+k)
			dfs(track)
			track = track[:len(track)-1]
		}

		if k > 0 && lastDigit-k >= 0 {
			track = append(track, lastDigit-k)
			dfs(track)
			track = track[:len(track)-1]
		}
	}

	for i := 1; i <= 9; i++ {
		dfs([]int{i})
	}

	return res
}

func list2Num(track []int) int {
	var sum int
	for i, num := range track {
		sum += int(float64(num) * math.Pow(10, float64(len(track)-1-i)))
	}
	return sum
}

//leetcode submit region end(Prohibit modification and deletion)
