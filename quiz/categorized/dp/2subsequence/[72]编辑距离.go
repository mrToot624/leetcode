package _subsequence

// leetcode submit region begin(Prohibit modification and deletion)
func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1) // minDistance of word1 and word2, when len(word1) = i, len(word2) = j
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
		dp[i][0] = i
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1] // skip
			} else {
				dp[i][j] = minOFThree(
					dp[i-1][j-1], // replace
					dp[i][j-1],   // insert
					dp[i-1][j],   // remove
				) + 1
			}
		}
	}

	return dp[m][n]
}

func minOFThree(a, b, c int) int {
	min := a
	if b < a {
		min = b
	}
	if c < min {
		min = c
	}
	return min
}

//leetcode submit region end(Prohibit modification and deletion)
