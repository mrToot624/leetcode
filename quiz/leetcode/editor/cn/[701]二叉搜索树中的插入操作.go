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
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}

	cur := root
	for cur != nil {
		if val < cur.Val {
			if cur.Left == nil {
				cur.Left = &TreeNode{Val: val}
				break
			}
			cur = cur.Left
		} else {
			if cur.Right == nil {
				cur.Right = &TreeNode{Val: val}
				break
			}
			cur = cur.Right
		}
	}
	return root
}
//leetcode submit region end(Prohibit modification and deletion)
