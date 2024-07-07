package main

import "math"

//leetcode submit region begin(Prohibit modification and deletion)
func minWindow(s string, t string) string {
	left, right := 0, 0

	targetCountMap, window := make(map[byte]int), make(map[byte]int)
	for i := 0; i < len(t); i++ {
		targetCountMap[t[i]]++
	}

	start, end := 0, math.MaxInt
	var satisfiedRunes int
	for right < len(s) {
		toAdd := s[right]
		if _, ok := targetCountMap[toAdd]; ok {
			window[toAdd]++
			if window[toAdd] == targetCountMap[toAdd] {
				satisfiedRunes++
			}
		}
		right++

		for satisfiedRunes == len(targetCountMap) && left < right {
			if right-left < end-start {
				start = left
				end = right
			}

			toRemove := s[left]
			if _, ok := targetCountMap[toRemove]; ok {
				if window[toRemove] == targetCountMap[toRemove] {
					satisfiedRunes--
				}
				window[toRemove]--
			}
			left++
		}
	}

	if end == math.MaxInt {
		end = 0
	}
	return s[start: end]
}
//leetcode submit region end(Prohibit modification and deletion)
