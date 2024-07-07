package main

// leetcode submit region begin(Prohibit modification and deletion)
func sortArray(nums []int) []int {
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

	merge(nums, tmp, start, end)
}

func merge(nums, tmp []int, start, end int) {
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
