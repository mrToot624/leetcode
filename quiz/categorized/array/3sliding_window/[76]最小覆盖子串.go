package main

import (
	"math"
)

//leetcode submit region begin(Prohibit modification and deletion)
func minWindow(s string, t string) string {
	left, right := 0, 0

	targetCountMap, window := make(map[rune]int), make(map[rune]int)
	for _, c := range t {
		targetCountMap[c]++
	}

	start, end := 0, math.MaxInt
	satisfiedRunes := make(map[rune]struct{})
	runes := []rune(s)
	for right < len(s) {
		if _, ok := targetCountMap[runes[right]]; ok {
			window[runes[right]]++
		}
		if window[runes[right]] >= targetCountMap[runes[right]] {
			satisfiedRunes[runes[right]] = struct{}{}
		}
		right++

		for len(satisfiedRunes) == len(targetCountMap) && left < right {
			if right-left < end-start {
				start = left
				end = right
			}

			if _, ok := targetCountMap[runes[left]]; ok {
				window[runes[left]]--
			}
			if window[runes[left]] < targetCountMap[runes[left]] {
				delete(satisfiedRunes, runes[left])
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