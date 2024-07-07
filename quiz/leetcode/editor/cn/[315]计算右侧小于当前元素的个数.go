package main

// leetcode submit region begin(Prohibit modification and deletion)
type pair struct {
	num   int
	index int
}

func countSmaller(nums []int) []int {
	count := make([]int, len(nums))

	pairs := make([]pair, len(nums))
	tmp := make([]pair, len(nums))
	for i, num := range nums {
		pairs[i] = pair{num, i}
	}

	countSmallerByMergeSort(pairs, tmp, count, 0, len(nums)-1)

	return count
}

func countSmallerByMergeSort(pairs, tmp []pair, count []int, start, end int) {
	if start == end {
		return
	}

	mid := start + (end-start)/2

	countSmallerByMergeSort(pairs, tmp, count, start, mid)
	countSmallerByMergeSort(pairs, tmp, count, mid+1, end)

	mergeAndCount(pairs, tmp, count, start, end)
}

func mergeAndCount(pairs, tmp []pair, count []int, start, end int) {
	for i := start; i <= end; i++ {
		tmp[i] = pairs[i]
	}

	mid := start + (end-start)/2

	left, right := start, mid+1
	for k := start; k <= end; k++ {
		if right == end+1 {
			count[tmp[left].index] += right - (mid + 1)
			pairs[k] = tmp[left]
			left++
		} else if left == mid+1 {
			pairs[k] = tmp[right]
			right++
		} else if tmp[left].num <= tmp[right].num {
			count[tmp[left].index] += right - (mid + 1)
			pairs[k] = tmp[left]
			left++
		} else {
			pairs[k] = tmp[right]
			right++
		}
	}
}

//leetcode submit region end(Prohibit modification and deletion)
