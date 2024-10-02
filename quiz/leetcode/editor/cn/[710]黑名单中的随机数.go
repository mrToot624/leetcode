package main

import "math/rand"

// leetcode submit region begin(Prohibit modification and deletion)
type Solution struct {
	sep   int
	remap map[int]int
}

func Constructor(n int, blacklist []int) Solution {
	blackSet := make(map[int]bool)
	for _, i := range blacklist {
		blackSet[i] = true
	}

	whiteNum, sep := n-1, n-len(blacklist)
	remap := make(map[int]int)
	for _, blackNum := range blacklist {
		if blackNum >= sep {
			continue
		}

		for blackSet[whiteNum] {
			whiteNum--
		}
		remap[blackNum] = whiteNum
		whiteNum--
	}
	return Solution{sep, remap}
}

func (this *Solution) Pick() int {
	num := rand.Intn(this.sep)
	if whiteNum, ok := this.remap[num]; ok {
		return whiteNum
	}
	return num
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(n, blacklist);
 * param_1 := obj.Pick();
 */
//leetcode submit region end(Prohibit modification and deletion)
