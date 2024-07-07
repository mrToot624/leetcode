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

var diameter int

func diameterOfBinaryTree(root *TreeNode) int {
	// have to add this to forcibly remove results from last test case
	// it seems that these global values will also become global in leetcode's submission system,
	// not per test scope
	diameter = 0

	maxDepthWithDiameter(root)
	return diameter
}

func maxDepthWithDiameter(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := maxDepthWithDiameter(root.Left)
	rightDepth := maxDepthWithDiameter(root.Right)

	diameter = max_543(leftDepth+rightDepth, diameter)

	return max_543(leftDepth, rightDepth)+1
}

func max_543(a, b int) int {
	if a < b {
		return b
	}
	return a
}
//leetcode submit region end(Prohibit modification and deletion)
