package main

// leetcode submit region begin(Prohibit modification and deletion)
func sortedSquares(nums []int) []int {
	res := make([]int, len(nums))
	left, right := 0, len(nums)-1
	for i := len(nums)-1; i >= 0; i-- {
		leftSquare, rightSquare := nums[left]*nums[left], nums[right]*nums[right]
		if leftSquare > rightSquare {
			res[i] = leftSquare
			left++
		} else {
			res[i] = rightSquare
			right--
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
