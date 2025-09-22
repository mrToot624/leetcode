package main

import "unicode"

// leetcode submit region begin(Prohibit modification and deletion)
func calculate(s string) int {
	op := '+'
	var sum int
	var opBeforeP []byte
	for i := 0; i < len(s); {
		switch s[i] {
		case ' ':
			i++
		case '+':
			if len(opBeforeP) > 0 && opBeforeP[len(opBeforeP)-1] == '-' {
				op = '-'
			} else {
				op = '+'
			}
			i++
		case '-':
			if len(opBeforeP) > 0 && opBeforeP[len(opBeforeP)-1] == '-' {
				op = '+'
			} else {
				op = '-'
			}
			i++
		case '(':
			opBeforeP = append(opBeforeP, byte(op))
			i++
		case ')':
			opBeforeP = opBeforeP[:len(opBeforeP)-1]
			i++
		default:
			var num int
			for i < len(s) && unicode.IsDigit(rune(s[i])) {
				digit := s[i] - '0'
				num = num*10 + int(digit)
				i++
			}
			if op == '+' {
				sum += num
			} else {
				sum -= num
			}
		}
	}
	return sum
}

//leetcode submit region end(Prohibit modification and deletion)
