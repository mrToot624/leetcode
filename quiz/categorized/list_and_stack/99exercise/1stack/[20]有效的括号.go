package main

// leetcode submit region begin(Prohibit modification and deletion)
func isValid(s string) bool {
	mapper := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	var stack []rune
	for _, c := range s {
		if c == '(' || c == '[' || c == '{' {
			stack = append(stack, c)
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != mapper[c] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

//leetcode submit region end(Prohibit modification and deletion)
