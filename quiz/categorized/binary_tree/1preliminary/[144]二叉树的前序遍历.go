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

var preOrderRes []int

func preorderTraversal(root *TreeNode) []int {
	// have to add this to forcibly remove results from last test case
	// it seems that these global values will also become global in leetcode's submission system, not per test scope
	preOrderRes = nil

	traverse_144(root)
	return preOrderRes
}

func traverse_144(root *TreeNode) {
	if root == nil {
		return
	}

	preOrderRes = append(preOrderRes, root.Val)
	traverse_144(root.Left)
	traverse_144(root.Right)
}
//leetcode submit region end(Prohibit modification and deletion)
