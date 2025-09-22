package main

// leetcode submit region begin(Prohibit modification and deletion)
func shipWithinDays(weights []int, days int) int {
	maxW, sum := weights[0], weights[0]
	for i := 1; i < len(weights); i++ {
		maxW = max(maxW, weights[i])
		sum += weights[i]
	}

	left, right := maxW, sum
	for left <= right {
		mid := left + (right-left)/2
		if canShipWithinDays(days, mid, weights) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return right + 1
}

func canShipWithinDays(days int, beltW int, weights []int) bool {
	var dayCnt, wSum int
	for _, weight := range weights {
		wSum += weight
		if wSum > beltW {
			wSum = weight
			dayCnt++
		}
	}
	return dayCnt+1 <= days
}

//leetcode submit region end(Prohibit modification and deletion)
