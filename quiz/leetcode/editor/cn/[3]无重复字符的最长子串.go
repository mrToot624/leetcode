package main

//leetcode submit region begin(Prohibit modification and deletion)

//func lengthOfLongestSubstring(s string) int {
//	var res, left, right int
//	window := make(map[byte]int)
//
//	for right < len(s) {
//		toAdd := s[right]
//		window[toAdd]++
//		right++
//
//		for window[toAdd] > 1 {
//			window[s[left]]--
//			left++
//		}
//
//		res = max_3(res, right-left)
//	}
//	return res
//}

func lengthOfLongestSubstring(s string) int {
	var res, left, right int
	window := make(map[byte]struct{})

	for right < len(s) {
		toAdd := s[right]
		_, found := window[toAdd]
		if !found {
			window[toAdd] = struct{}{}
		} else {
			res = max_3(res, right-left)
			for s[left] != toAdd {
				delete(window, s[left])
				left++
			}
			left++
		}
		right++
	}
	return max_3(res, right-left)
}

func max_3(a, b int) int {
	if a < b {
		return b
	}
	return a
}

//leetcode submit region end(Prohibit modification and deletion)
