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
func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
	if root.Val < val {
		return &TreeNode{Val: val, Left: root}
	}
	cur := root
	for cur != nil {
		if cur.Right == nil {
			cur.Right = &TreeNode{Val: val}
			break
		} else if cur.Right.Val < val {
			cur.Right = &TreeNode{Val: val, Left: cur.Right}
			break
		}
		cur = cur.Right
	}
	return root
}

//leetcode submit region end(Prohibit modification and deletion)
