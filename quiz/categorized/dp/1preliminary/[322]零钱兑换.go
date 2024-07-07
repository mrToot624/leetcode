package main

//leetcode submit region begin(Prohibit modification and deletion)
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0

	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
		for _, coin := range coins {
			if i-coin < 0 {
				continue
			}
			dp[i] = min_322(dp[i], dp[i-coin]+1)
		}
	}

	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}

func recursiveCoinChangeWithCache(coins []int, amount int) int {
	cache := make(map[int]int)
	return coinChangeWithCache(coins, amount, cache)
}

func coinChangeWithCache(coins []int, amount int, cache map[int]int) int {
	if amount == 0 {
		return 0
	}

	if amount < 0 {
		return -1
	}

	if value, ok := cache[amount]; ok {
		return value
	}

	minCount := amount + 1
	for _, coin := range coins {
		count := coinChangeWithCache(coins, amount-coin, cache)
		if count == -1 {
			continue
		}
		minCount = min_322(minCount, count+1)
	}

	if minCount == amount + 1 {
		cache[amount] = -1
	} else {
		cache[amount] = minCount
	}

	return cache[amount]
}

func min_322(a, b int) int {
	if a < b {
		return a
	}
	return b
}
//leetcode submit region end(Prohibit modification and deletion)
