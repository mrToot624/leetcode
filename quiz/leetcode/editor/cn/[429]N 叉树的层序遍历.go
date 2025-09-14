package main

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func levelOrder(root *Node) [][]int {
    var res [][]int
	if root == nil {
		return res
	}

	q := []*Node{root}
	for len(q) > 0 {
		l := len(q)
		level := make([]int, l)
		for i := 0; i < l; i++ {
			top := q[0]
			q = q[1:]
			level[i] = top.Val
			for _, child := range top.Children {
				q = append(q, child)
			}
		}
		res = append(res, level)
	}
	return res
}
//leetcode submit region end(Prohibit modification and deletion)
