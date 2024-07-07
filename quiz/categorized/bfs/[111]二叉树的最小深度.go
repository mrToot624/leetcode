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
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	depth := 1
	var q []*TreeNode
	q = append(q, root)

	for len(q) > 0 {
		l := len(q)
		for i := 0; i < l; i++ {
			node := q[0]
			if node.Left == nil && node.Right == nil {
				return depth
			}

			if node.Left != nil {
				q = append(q, node.Left)
			}

			if node.Right != nil {
				q = append(q, node.Right)
			}
			q = q[1:]
		}
		depth++
	}

	return depth
}
//leetcode submit region end(Prohibit modification and deletion)
