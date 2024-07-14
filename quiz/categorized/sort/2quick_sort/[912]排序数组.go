package main

import "math/rand"

// leetcode submit region begin(Prohibit modification and deletion)
func sortArray(nums []int) []int {
	rand.Shuffle(len(nums), func(i, j int) {
		nums[i], nums[j] = nums[j], nums[i]
	})
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func quickSort(nums []int, start, end int) {
	if start >= end {
		return
	}

	p := quickSortPartition(nums, start, end)

	quickSort(nums, start, p-1)
	quickSort(nums, p+1, end)
}

func quickSortPartition(nums []int, start, end int) int {
	pivot := nums[start]

	i, j := start+1, end
	for i <= j {
		for i <= end && nums[i] <= pivot {
			i++
		}
		for nums[j] > pivot {
			j--
		}

		if i >= j {
			break
		}
		quickSortSwap(nums, i, j)
	}

	quickSortSwap(nums, start, j)
	return j
}

func quickSortSwap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

//leetcode submit region end(Prohibit modification and deletion)
