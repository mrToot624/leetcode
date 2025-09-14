package main

// leetcode submit region begin(Prohibit modification and deletion)
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	res := make([]int, len(nums1))
	tmp := make(map[int]int)

	var stack []int
	for i := len(nums2) - 1; i >= 0; i-- {
		for len(stack) > 0 && stack[len(stack)-1] <= nums2[i] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) == 0 {
			tmp[nums2[i]] = -1
		} else {
			tmp[nums2[i]] = stack[len(stack)-1]
		}
		stack = append(stack, nums2[i])
	}

	for i, num := range nums1 {
		res[i] = tmp[num]
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
