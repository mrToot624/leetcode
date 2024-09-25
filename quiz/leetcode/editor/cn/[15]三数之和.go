package main

import "sort"

// leetcode submit region begin(Prohibit modification and deletion)
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int

	target := 0
	for i, num := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		if twoSums := findAllTwoSums(nums[i+1:], target-num, num); twoSums != nil {
			res = append(res, twoSums...)
		}
	}
	return res
}

func findAllTwoSums(nums []int, target, curNum int) [][]int {
	var res [][]int
	start, end := 0, len(nums)-1
	for start < end {
		curSum := nums[start] + nums[end]
		if curSum == target {
			res = append(res, []int{curNum, nums[start], nums[end]})
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
