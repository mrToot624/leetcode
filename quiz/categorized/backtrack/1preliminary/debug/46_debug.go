package main

// leetcode submit region begin(Prohibit modification and deletion)
func permute(nums []int) [][]int {
	var res [][]int
	var track []int
	used := make([]bool, len(nums))
	backtrack_46(nums, track, used, &res)
	return res
}

func backtrack_46(nums, track []int, used []bool, res *[][]int) {
	if len(track) == len(nums) {
		*res = append(*res, track)
		return
	}

	for i, num := range nums {
		if used[i] {
			continue
		}

		used[i] = true
		track = append(track, num)
		backtrack_46(nums, track, used, res)
		track = track[:len(track)-1]
		used[i] = false
	}
}

//func main() {
//	nums := []int{5, 4, 6}
//	fmt.Println(permute(nums))
//}
//leetcode submit region end(Prohibit modification and deletion)
