package main

//leetcode submit region begin(Prohibit modification and deletion)
func findAnagrams(s string, p string) []int {
	left, right := 0, 0
	target, window := make(map[rune]int), make(map[rune]int)
	for _, c := range p {
		target[c]++
	}

	runes := []rune(s)
	var satisfied int
	var res []int

	for right < len(s) {
		toAdd := runes[right]
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

			toRemove := runes[left]
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
