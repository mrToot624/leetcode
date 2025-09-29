package main

import "strconv"

// leetcode submit region begin(Prohibit modification and deletion)
func evalRPN(tokens []string) int {
	var stack []int
	calculateByOp := func(op string) {
		a, b := stack[len(stack)-2], stack[len(stack)-1]
		stack = stack[:len(stack)-2]
		var num int
		switch op {
		case "+":
			num = a + b
		case "-":
			num = a - b
		case "*":
			num = a * b
		case "/":
			num = a / b
		}
		stack = append(stack, num)
	}
	for _, token := range tokens {
		if token == "+" || token == "-" || token == "*" || token == "/" {
			calculateByOp(token)
		} else {
			num, _ := strconv.Atoi(token)
			stack = append(stack, num)
		}
	}
	return stack[len(stack)-1]
}

//leetcode submit region end(Prohibit modification and deletion)
