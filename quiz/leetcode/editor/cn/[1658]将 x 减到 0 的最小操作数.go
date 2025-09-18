package main

// leetcode submit region begin(Prohibit modification and deletion)
func minOperations(nums []int, x int) int {
	var sum int
	for _, num := range nums {
		sum += num
	}
	target := sum - x

	if target < 0 {
		return -1
	} else if target == 0 {
		return len(nums)
	}

	maxL, sum := 0, 0
	left, right := 0, 0
	for right < len(nums) {
		sum += nums[right]
		right++


		for sum > target && left < right {
			sum -= nums[left]
			left++
		}

		if sum == target {
			maxL = max(maxL, right-left)
		}
	}
	if maxL == 0 {
		return -1
	}
	return len(nums) - maxL
}

func minOperations_timeout(nums []int, x int) int {
	if len(nums) == 0 || x < 0 {
		return -1
	}
	if x == nums[0] || x == nums[len(nums)-1] {
		return 1
	}

	leftOp := minOperations_timeout(nums[1:], x-nums[0])
	rightOp := minOperations_timeout(nums[:len(nums)-1], x-nums[len(nums)-1])
	if leftOp == -1 && rightOp == -1 {
		return -1
	}
	if leftOp == -1 {
		return rightOp + 1
	} else if rightOp == -1 {
		return leftOp + 1
	}
	return min(leftOp, rightOp) + 1
}

//leetcode submit region end(Prohibit modification and deletion)
