package main

//leetcode submit region begin(Prohibit modification and deletion)
func checkInclusion(s1 string, s2 string) bool {
	left, right := 0, 0
	targetCount, window := make(map[rune]int), make(map[rune]int)

	for _, c := range s1 {
		targetCount[c]++
	}

	runes := []rune(s2)
	var satisfied int
	for right < len(s2) {
		toAdd := runes[right]
		if _, ok := targetCount[toAdd]; ok {
			window[toAdd]++
			if window[toAdd] == targetCount[toAdd] {
				satisfied++
			}
		}
		right++

		for right-left >= len(s1) {
			if satisfied == len(targetCount) {
				return true
			}

			toRemove := runes[left]
			if _, ok := targetCount[toRemove]; ok {
				window[toRemove]--
				if window[toRemove] < targetCount[toRemove] {
					satisfied--
				}
			}
			left++
		}
	}

	return false
}
//leetcode submit region end(Prohibit modification and deletion)
