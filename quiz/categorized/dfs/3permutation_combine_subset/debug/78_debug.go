package main

import "fmt"

// leetcode submit region begin(Prohibit modification and deletion)
func subsets(nums []int) [][]int {
	var res [][]int
	var track []int
	backtrack_78(track, nums, &res, 0)
	return res
}

func backtrack_78(track, nums []int, res *[][]int, start int) {
	subset := make([]int, len(track))
	copy(subset, track)
	*res = append(*res, subset)

	for i := start; i < len(nums); i++ {
		track = append(track, nums[i])
		backtrack_78(track, nums, res, start+1)
		track = track[:len(track)-1]
	}
}

func main() {
	test := []int{1, 2, 3}
	fmt.Println(subsets(test))
}

//leetcode submit region end(Prohibit modification and deletion)
