package main

import "math/rand"

// leetcode submit region begin(Prohibit modification and deletion)
type Solution struct {
	targetIndex map[int][]int
}

func Constructor(nums []int) Solution {
	targetIndexMap := make(map[int][]int)
	for i, num := range nums {
		targetIndexMap[num] = append(targetIndexMap[num], i)
	}
	return Solution{targetIndexMap}
}

func (this *Solution) Pick(target int) int {
	return this.targetIndex[target][rand.Intn(len(this.targetIndex[target]))]
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Pick(target);
 */
//leetcode submit region end(Prohibit modification and deletion)
