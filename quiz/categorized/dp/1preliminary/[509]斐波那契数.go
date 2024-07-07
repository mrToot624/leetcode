package main

//leetcode submit region begin(Prohibit modification and deletion)
func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	prev, cur := 0, 1
	var sum int
	for i := 2; i <= n; i++ {
		sum = prev + cur
		prev = cur
		cur = sum
	}
	return sum
}

func fibWithRedundantDP(n int) int {
	dp := make([]int, n+1)
	if n > 0 {
		dp[1] = 1
	}

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

func RecursiveFibWithCache(n int) int {
	cache := make([]int, n+1)
	return fibWithCache(n, cache)
}

func fibWithCache(n int, cache []int) (res int) {
	defer func() {
		cache[n] = res
	}()

	if n == 0 || n == 1 {
		res = n
		return
	}

	if cache[n] != 0 {
		return cache[n]
	}

	res = fibWithCache(n-1, cache) + fibWithCache(n-2, cache)
	return
}
//leetcode submit region end(Prohibit modification and deletion)
