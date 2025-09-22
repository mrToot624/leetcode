package main

import "unicode"

// leetcode submit region begin(Prohibit modification and deletion)
func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		for left < right && !isLetterOrNum(rune(s[left])) {
			left++
		}
		for left < right && !isLetterOrNum(rune(s[right])) {
			right--
		}
		if left < right && unicode.ToLower(rune(s[left])) != unicode.ToLower(rune(s[right])) {
			return false
		}
		left++
		right--
	}
	return true
}

func isLetterOrNum(r rune) bool {
	return (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9')
}

//leetcode submit region end(Prohibit modification and deletion)
