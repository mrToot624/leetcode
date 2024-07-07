package main

// leetcode submit region begin(Prohibit modification and deletion)
func reversePairs(nums []int) int {
	var count int
	tmp := make([]int, len(nums))

	countReversePairsByMergeSort(nums, tmp, &count, 0, len(nums)-1)

	return count
}

func countReversePairsByMergeSort(nums, tmp []int, count *int, start, end int) {
	if start == end {
		return
	}

	mid := start + (end-start)/2

	countReversePairsByMergeSort(nums, tmp, count, start, mid)
	countReversePairsByMergeSort(nums, tmp, count, mid+1, end)

	mergeAndCountReversePairs(nums, tmp, count, start, end)
}

func mergeAndCountReversePairs(nums, tmp []int, count *int, start, end int) {
	for i := start; i <= end; i++ {
		tmp[i] = nums[i]
	}

	mid := start + (end-start)/2

	right := mid + 1
	for left := start; left <= mid; left++ {
		for right <= end && nums[left] > 2*nums[right] {
			right++
		}
		*count += right - (mid + 1)
	}

	left, right := start, mid+1
	for k := start; k <= end; k++ {
		if left == mid+1 {
			nums[k] = tmp[right]
			right++
		} else if right == end+1 {
			nums[k] = tmp[left]
			left++
		} else if tmp[left] <= tmp[right] {
			nums[k] = tmp[left]
			left++
		} else {
			nums[k] = tmp[right]
			right++
		}
	}
}

//leetcode submit region end(Prohibit modification and deletion)
