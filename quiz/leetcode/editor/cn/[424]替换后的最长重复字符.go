package main

// leetcode submit region begin(Prohibit modification and deletion)
func characterReplacement(s string, k int) int {
	charCnt := make(map[byte]int)
	left, right := 0, 0
	maxCnt := 0

	var res int
	for right < len(s) {
		charCnt[s[right]]++
		maxCnt = max(maxCnt, charCnt[s[right]])
		right++
		if right-left-maxCnt > k {
			charCnt[s[left]]--
			left++
		}
		res = max(res, right-left)
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
