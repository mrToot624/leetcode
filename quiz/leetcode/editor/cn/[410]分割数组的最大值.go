package main

import "sort"

// leetcode submit region begin(Prohibit modification and deletion)
func splitArray(nums []int, k int) int {
	left, right := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		left = max(left, nums[i])
		right += nums[i]
	}

	return left + sort.Search(right-left, func(max int) bool {
		return canSplitWithMax(nums, k, max+left)
	})

	//for left <= right {
	//	mid := left + (right-left)/2
	//	if canSplitWithMax(nums, k, mid) {
	//		right = mid - 1
	//	} else {
	//		left = mid + 1
	//	}
	//}
	//return right + 1
}

func canSplitWithMax(nums []int, k, max int) bool {
	sum, cnt := 0, 1
	for _, num := range nums {
		sum += num
		if sum > max {
			cnt++
			sum = num
		}
	}
	return cnt <= k
}

//leetcode submit region end(Prohibit modification and deletion)
