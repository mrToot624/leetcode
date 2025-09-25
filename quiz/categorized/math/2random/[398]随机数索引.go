package main

import "math/rand"

// leetcode submit region begin(Prohibit modification and deletion)
type Solution_398 struct {
	targetIndex map[int][]int
}

func Constructor_398(nums []int) Solution_398 {
	targetIndexMap := make(map[int][]int)
	for i, num := range nums {
		targetIndexMap[num] = append(targetIndexMap[num], i)
	}
	return Solution_398{targetIndexMap}
}

func (this *Solution_398) Pick(target int) int {
	return this.targetIndex[target][rand.Intn(len(this.targetIndex[target]))]
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Pick(target);
 */
//leetcode submit region end(Prohibit modification and deletion)
