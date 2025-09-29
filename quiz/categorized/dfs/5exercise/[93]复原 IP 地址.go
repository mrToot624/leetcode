package main

import (
	"strconv"
	"strings"
)

// leetcode submit region begin(Prohibit modification and deletion)
func restoreIpAddresses(s string) []string {
	n := len(s)
	var track, res []string
	var dfs func(int)
	dfs = func(start int) {
		if start == n {
			if len(track) == 4 {
				res = append(res, strings.Join(track, "."))
			}
			return
		}

		if n-start > (4-len(track))*3 {
			return
		}

		for i := start; i < n && i < start+3; i++ {
			if i > start && !isValidIP(s[start:i+1]) {
				continue
			}
			track = append(track, s[start:i+1])
			dfs(i + 1)
			track = track[:len(track)-1]
		}
	}
	dfs(0)
	return res
}

func isValidIP(s string) bool {
	if len(s) >= 2 && s[0] == '0' {
		return false
	}
	num, _ := strconv.Atoi(s)
	return num <= 255
}

//leetcode submit region end(Prohibit modification and deletion)
