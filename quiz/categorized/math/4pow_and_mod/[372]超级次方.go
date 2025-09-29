package main

// leetcode submit region begin(Prohibit modification and deletion)
const mod = 1337

func superPow(a int, b []int) int {
	res := 1
	for i := len(b) - 1; i >= 0; i-- {
		res = (res * quickPowWithMod(a, b[i])) % mod
		a = quickPowWithMod(a, 10)
	}
	return res
}

func quickPowWithMod(a int, b int) int {
	if b == 0 {
		return 1
	}
	if b&1 == 1 {
		return (a * quickPowWithMod(a, b-1)) % mod
	}
	half := quickPowWithMod(a, b>>1) % mod
	return (half * half) % mod
}

//leetcode submit region end(Prohibit modification and deletion)
