package main

// leetcode submit region begin(Prohibit modification and deletion)
func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	for i, num := range nums {
		if another, ok := numMap[target-num]; ok {
			return []int{i, another}
		} else {
			numMap[num] = i
		}
	}
	return nil
}

//leetcode submit region end(Prohibit modification and deletion)
