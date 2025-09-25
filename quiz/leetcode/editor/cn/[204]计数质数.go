package main

//leetcode submit region begin(Prohibit modification and deletion)
func countPrimes(n int) int {
	var cnt int
    isPrime := make([]int, n)
	for i := 2; i < n; i++ {
		if isPrime[i] == 0 {
			cnt++
			for j := i * i; j < n; j += i {
				isPrime[j] = 1
			}
		}
	}
	return cnt
}
//leetcode submit region end(Prohibit modification and deletion)
