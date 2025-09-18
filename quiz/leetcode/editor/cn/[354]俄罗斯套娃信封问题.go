package main

import "sort"

// leetcode submit region begin(Prohibit modification and deletion)
func maxEnvelopes(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] > envelopes[j][1]
		}
		return envelopes[i][0] < envelopes[j][0]
	})

	var piles []int
	for _, envelope := range envelopes {
		left, right := 0, len(piles)-1
		for left <= right {
			mid := left + (right-left)>>1
			if envelope[1] <= piles[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
		if left == len(piles) {
			piles = append(piles, envelope[1])
		} else {
			piles[left] = envelope[1]
		}
	}
	return len(piles)
}

//leetcode submit region end(Prohibit modification and deletion)

func maxEnvelopes_timeout(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		return envelopes[i][0] < envelopes[j][0]
	})

	dp := make([]int, len(envelopes))
	dp[0] = 1
	for i := 1; i < len(envelopes); i++ {
		maxL := 0
		for j := 0; j < i; j++ {
			if envelopes[j][0] < envelopes[i][0] && envelopes[j][1] < envelopes[i][1] {
				maxL = max(maxL, dp[j])
			}
		}
		dp[i] = maxL + 1
	}
	maxL := dp[0]
	for i := 1; i < len(dp); i++ {
		maxL = max(maxL, dp[i])
	}
	return maxL
}
