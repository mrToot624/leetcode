package main

//leetcode submit region begin(Prohibit modification and deletion)
func longestPalindrome(s string) string {
	var res string

	for i := 0; i < len(s); i++ {
		odd := findLongestPalindromeByIndex(s, i, i)
		even := findLongestPalindromeByIndex(s, i, i+1)

		longer := even
		if len(longer) < len(odd) {
			longer = odd
		}

		if len(res) < len(longer) {
			res = longer
		}
	}
	return res
}

func findLongestPalindromeByIndex(s string, left, right int) string {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}

	return s[left+1: right]
}
//leetcode submit region end(Prohibit modification and deletion

