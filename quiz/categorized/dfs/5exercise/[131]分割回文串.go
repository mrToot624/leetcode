package main

// leetcode submit region begin(Prohibit modification and deletion)
func partition(s string) [][]string {
	n := len(s)
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
		dp[i][i] = true
	}
	for l := 1; l < n; l++ {
		for i := 0; i+l < n; i++ {
			j := i + l
			if l == 1 {
				dp[i][j] = s[i] == s[j]
			} else {
				dp[i][j] = dp[i+1][j-1] && s[i] == s[j]
			}
		}
	}

	var res [][]string
	var track []string
	var dfs func(i int)
	dfs = func(i int) {
		if i == n {
			tmp := make([]string, len(track))
			copy(tmp, track)
			res = append(res, tmp)
			return
		}

		for j := i; j < n; j++ {
			if dp[i][j] {
				track = append(track, s[i:j+1])
				dfs(j + 1)
				track = track[:len(track)-1]
			}
		}
	}
	dfs(0)
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
