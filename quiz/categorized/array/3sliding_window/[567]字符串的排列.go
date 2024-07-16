package main

//leetcode submit region begin(Prohibit modification and deletion)
func checkInclusion(s1 string, s2 string) bool {
	left, right := 0, 0
	targetCount, window := make(map[byte]int), make(map[byte]int)
	for i := 0; i < len(s1); i++ {
		targetCount[s1[i]]++
	}

	var satisfied int
	for right < len(s2) {
		toAdd := s2[right]
		if _, ok := targetCount[toAdd]; ok {
			window[toAdd]++
			if window[toAdd] == targetCount[toAdd] {
				satisfied++
			}
		}
		right++

		if right-left == len(s1) {
			if satisfied == len(targetCount) {
				return true
			}

			toRemove := s2[left]
			if _, ok := targetCount[toRemove]; ok {
				if window[toRemove] == targetCount[toRemove] {
					satisfied--
				}
				window[toRemove]--
			}
			left++
		}
	}

	return false
}
//leetcode submit region end(Prohibit modification and deletion)
