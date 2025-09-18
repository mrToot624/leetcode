package main

// leetcode submit region begin(Prohibit modification and deletion)
func minEatingSpeed(piles []int, h int) int {
	var sum, right int
	for _, pile := range piles {
		right = max(right, pile)
		sum += pile
	}
	left := max(1, sum/h)

	for left <= right {
		mid := left + (right-left)/2
		if canEatUp(mid, h, piles) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return right + 1
}

func canEatUp(speed, h int, piles []int) bool {
	var sum int
	for _, pile := range piles {
		t := pile / speed
		if pile%speed != 0 {
			t++
		}
		sum += t
	}
	return sum <= h
}

//leetcode submit region end(Prohibit modification and deletion)
