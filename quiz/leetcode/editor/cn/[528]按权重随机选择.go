package main

import (
	"math/rand"
	"sort"
)

// leetcode submit region begin(Prohibit modification and deletion)
type Solution struct {
	preSum []int
}

func Constructor(w []int) Solution {
	preSum := make([]int, len(w)+1)
	for i := 0; i < len(w); i++ {
		preSum[i+1] = preSum[i] + w[i]
	}
	return Solution{preSum}
}

func (this *Solution) PickIndex() int {
	//x := rand.Intn(this.preSum[len(this.preSum)-1]) + 1
	//left, right := 1, len(this.preSum)-1
	//for left <= right {
	//	mid := left + (right-left)>>1
	//	if this.preSum[mid] < x {
	//		left = mid + 1
	//	} else {
	//		right = mid - 1
	//	}
	//}
	//return right
	return sort.SearchInts(this.preSum, rand.Intn(this.preSum[len(this.preSum)-1])+1) - 1
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */
//leetcode submit region end(Prohibit modification and deletion)
