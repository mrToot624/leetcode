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

// merge sort
func sortArrayByMergeSort(nums []int) []int {
	tmp := make([]int, len(nums))
	mergeSort(nums, tmp, 0, len(nums)-1)
	return nums
}

func mergeSort(nums, tmp []int, start, end int) {
	if start >= end {
		return
	}

	mid := start + (end-start)/2

	mergeSort(nums, tmp, start, mid)
	mergeSort(nums, tmp, mid+1, end)

	merge_912(nums, tmp, start, end)
}

func merge_912(nums, tmp []int, start, end int) {
	for i := start; i <= end; i++ {
		tmp[i] = nums[i]
	}
	mid := start + (end-start)/2

	i, j := start, mid+1
	for k := start; k <= end; k++ {
		if i == mid+1 {
			nums[k] = tmp[j]
			j++
		} else if j == end+1 {
			nums[k] = tmp[i]
			i++
		} else if tmp[i] <= tmp[j] {
			nums[k] = tmp[i]
			i++
		} else {
			nums[k] = tmp[j]
			j++
		}
	}
}

//leetcode submit region end(Prohibit modification and deletion)
