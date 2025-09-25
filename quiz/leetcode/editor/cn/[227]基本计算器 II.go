package main

// leetcode submit region begin(Prohibit modification and deletion)
func calculate(s string) int {
	var stack []int
	prevOp := 1
	for i := 0; i < len(s); {
		switch s[i] {
		case ' ':
			i++
		case '+':
			prevOp = 1
			i++
		case '-':
			prevOp = -1
			i++
		case '*', '/':
			op := s[i]
			if op == '*' {
				stack[len(stack)-1] *= getNextNum(s, &i)
			} else {
				stack[len(stack)-1] /= getNextNum(s, &i)
			}
		default:
			stack = append(stack, prevOp*getNextNum(s, &i))
		}
	}

	sum := 0
	for _, num := range stack {
		sum += num
	}
	return sum
}

func getNextNum(s string, i *int) int {
	num := 0
	for !(s[*i] >= '0' && s[*i] <= '9') {
		*i++
	}
	for *i < len(s) && s[*i] >= '0' && s[*i] <= '9' {
		digit := s[*i] - '0'
		num = num*10 + int(digit)
		*i++
	}
	return num
}

//leetcode submit region end(Prohibit modification and deletion)
