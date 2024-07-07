package main

// leetcode submit region begin(Prohibit modification and deletion)
func countRangeSum(nums []int, lower int, upper int) int {
	preSum := make([]int, len(nums)+1)
	for i, num := range nums {
		preSum[i+1] = preSum[i] + num
	}

	tmp := make([]int, len(preSum))
	var count int

	countRangeSumByMergeSort(preSum, tmp, &count, lower, upper, 0, len(preSum)-1)

	return count
}

func countRangeSumByMergeSort(preSum, tmp []int, count *int, lower, upper, start, end int) {
	if start == end {
		return
	}

	mid := start + (end-start)/2

	countRangeSumByMergeSort(preSum, tmp, count, lower, upper, start, mid)
	countRangeSumByMergeSort(preSum, tmp, count, lower, upper, mid+1, end)

	mergeAndCountRangeSum(preSum, tmp, count, lower, upper, start, end)
}

func mergeAndCountRangeSum(preSum, tmp []int, count *int, lower, upper, start, end int) {
	for i := start; i <= end; i++ {
		tmp[i] = preSum[i]
	}

	mid := start + (end-start)/2

	rightL, rightU := mid+1, mid+1
	for left := start; left <= mid; left++ {
		for rightL <= end && preSum[rightL] - preSum[left] < lower {
			rightL++
		}
		for rightU <= end && preSum[rightU] - preSum[left] <= upper {
			rightU++
		}
		*count += rightU - rightL
	}

	left, right := start, mid+1
	for k := start; k <= end; k++ {
		if left == mid+1 {
			preSum[k] = tmp[right]
			right++
		} else if right == end+1 {
			preSum[k] = tmp[left]
			left++
		} else if tmp[left] <= tmp[right] {
			preSum[k] = tmp[left]
			left++
		} else {
			preSum[k] = tmp[right]
			right++
		}
	}
}

//leetcode submit region end(Prohibit modification and deletion)
