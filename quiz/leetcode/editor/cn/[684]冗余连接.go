package main

// leetcode submit region begin(Prohibit modification and deletion)
func findRedundantConnection(edges [][]int) []int {
	n := len(edges)
	uf := NewUnionFind(n + 1)
	var res []int
	for _, edge := range edges {
		p, q := edge[0], edge[1]
		if uf.Connected(p, q) {
			res = edge
			break
		} else {
			uf.Union(p, q)
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
