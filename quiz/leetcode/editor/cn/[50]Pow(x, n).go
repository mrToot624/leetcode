package main

// leetcode submit region begin(Prohibit modification and deletion)
func myPow(x float64, n int) float64 {
	if n >= 0 {
		return quickPow(x, n)
	}
	return 1.0 / quickPow(x, -n)
}

func quickPow(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	}
	if n&1 == 1 {
		return x * quickPow(x, n-1)
	}
	half := quickPow(x, n>>1)
	return half * half
}

//leetcode submit region end(Prohibit modification and deletion)
