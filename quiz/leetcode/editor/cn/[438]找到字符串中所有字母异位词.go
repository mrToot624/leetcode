package main

//leetcode submit region begin(Prohibit modification and deletion)
func findAnagrams(s string, p string) []int {
	left, right := 0, 0
	target, window := make(map[byte]int), make(map[byte]int)
	for i := 0; i < len(p); i++ {
		target[p[i]]++
	}

	var satisfied int
	var res []int

	for right < len(s) {
		toAdd := s[right]
		if _, ok := target[toAdd]; ok {
			window[toAdd]++
			if window[toAdd] == target[toAdd] {
				satisfied++
			}
		}
		right++

		for right-left >= len(p) {
			if satisfied == len(target) {
				res = append(res, left)
			}

			toRemove := s[left]
			if _, ok := target[toRemove]; ok {
				if window[toRemove] == target[toRemove] {
					satisfied--
				}
				window[toRemove]--
			}
			left++
		}
	}
	return res
}
//leetcode submit region end(Prohibit modification and deletion)
