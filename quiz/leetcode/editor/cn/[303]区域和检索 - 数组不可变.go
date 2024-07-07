package main

//leetcode submit region begin(Prohibit modification and deletion)
type NumArray struct {
	preSum []int
}


func Constructor_303(nums []int) NumArray {
	na := NumArray{preSum: make([]int, len(nums)+1)}
	for i, num := range nums {
		na.preSum[i+1] = na.preSum[i] + num
	}
	return na
}


func (this *NumArray) SumRange(left int, right int) int {
	return this.preSum[right+1] - this.preSum[left]
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
//leetcode submit region end(Prohibit modification and deletion)
