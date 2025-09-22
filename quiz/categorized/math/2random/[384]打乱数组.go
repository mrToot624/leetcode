package main

import "math/rand"

// leetcode submit region begin(Prohibit modification and deletion)
type Solution_384 struct {
	nums, original []int
}

func Constructor_384(nums []int) Solution_384 {
	return Solution_384{nums, append([]int(nil), nums...)}
}

func (this *Solution_384) Reset() []int {
	copy(this.nums, this.original)
	return this.nums
}

func (this *Solution_384) Shuffle() []int {
	n := len(this.nums)
	for i := 0; i < n; i++ {
		randI := i + rand.Intn(n-i)
		this.nums[i], this.nums[randI] = this.nums[randI], this.nums[i]
	}
	return this.nums
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
//leetcode submit region end(Prohibit modification and deletion)
