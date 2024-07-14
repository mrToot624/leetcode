package main

// leetcode submit region begin(Prohibit modification and deletion)
func findKthLargest(nums []int, k int) int {
	return quickSelectKthLargest(nums, k, 0, len(nums)-1)
}

func quickSelectKthLargest(nums []int, k, start, end int) int {
	if start == end {
		return nums[start]
	}

	p := quickSelectPartition(nums, start, end)
	if p == len(nums)-k {
		return nums[p]
	} else if p > len(nums)-k {
		return quickSelectKthLargest(nums, k, start, p-1)
	}
	return quickSelectKthLargest(nums, k, p+1, end)
}

func quickSelectPartition(nums []int, start, end int) int {
	pivot := nums[start]

	i, j := start+1, end
	for i <= j {
		for i <= end && nums[i] <= pivot {
			i++
		}
		for nums[j] > pivot {
			j--
		}

		if j <= i {
			break
		}
		quickSelectSwap(nums, i, j)
	}

	quickSelectSwap(nums, start, j)
	return j
}

func quickSelectSwap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

//leetcode submit region end(Prohibit modification and deletion)
