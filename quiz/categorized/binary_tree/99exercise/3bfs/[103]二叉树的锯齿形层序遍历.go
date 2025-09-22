package main

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	q := []*TreeNode{root}
	level := 1
	for len(q) > 0 {
		l := len(q)
		var vals []int
		for i := 0; i < l; i++ {
			cur := q[0]
			q = q[1:]
			if cur.Left != nil {
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			}
			if level%2 == 1 {
				vals = append(vals, cur.Val)
			} else {
				vals = append([]int{cur.Val}, vals...)
			}
		}
		res = append(res, vals)
		level++
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
