package main

import "sort"

// leetcode submit region begin(Prohibit modification and deletion)
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	return nSum(nums, target, 4)
}

func nSum(nums []int, target int, n int) [][]int {
	if n == 2 {
		return getTwoSum(nums, target)
	}

	var res [][]int
	for i, num := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for _, subSum := range nSum(nums[i+1:], target-num, n-1) {
			res = append(res, append(subSum, num))
		}
	}
	return res
}

func getTwoSum(nums []int, target int) [][]int {
	var res [][]int
	start, end := 0, len(nums)-1
	for start < end {
		curSum := nums[start] + nums[end]
		if curSum == target {
			res = append(res, []int{nums[start], nums[end]})
			curStart, curEnd := start, end
			for start < end && nums[start] == nums[curStart] {
				start++
			}
			for start < end && nums[end] == nums[curEnd] {
				end--
			}
		} else if curSum < target {
			start++
		} else {
			end--
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
