package main

// leetcode submit region begin(Prohibit modification and deletion)
func equationsPossible(equations []string) bool {
	uf := NewUnionFind(26)
	for _, equation := range equations {
		if p, q, isTrue := parseEquation(equation); isTrue {
			uf.Union(p, q)
		}
	}
	for _, equation := range equations {
		if p, q, isTrue := parseEquation(equation); !isTrue && uf.Connected(p, q) {
			return false
		}
	}
	return true
}

func parseEquation(equation string) (int, int, bool) {
	return int(equation[0] - 'a'), int(equation[3] - 'a'), equation[1] == '='
}

//leetcode submit region end(Prohibit modification and deletion)
