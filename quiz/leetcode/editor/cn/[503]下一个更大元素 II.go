package main

// leetcode submit region begin(Prohibit modification and deletion)
func nextGreaterElements(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	var stack []int

	for i := 2*n-1; i >= 0 ; i-- {
		realIndex := i%n
		for len(stack) > 0 && stack[len(stack)-1] <= nums[realIndex] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) == 0 {
			res[realIndex] = -1
		} else {
			res[realIndex] = stack[len(stack)-1]
		}
		stack = append(stack, nums[realIndex])
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
