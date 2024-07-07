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

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)

	return getDeeper(leftDepth, rightDepth)+1
}

func getDeeper(a, b int) int {
	if a < b {
		return b
	}
	return a
}

//leetcode submit region end(Prohibit modification and deletion)

var dep, maxDep int

func traverseMaxDepth(root *TreeNode) int {
	// have to add this to forcibly remove results from last test case
	// it seems that these global values will also become global in leetcode's submission system, not per test scope
	dep, maxDep = 0, 0

	traverse_104(root)
	return maxDep
}

func traverse_104(root *TreeNode) {
	if root == nil {
		return
	}

	dep++
	maxDep = getDeeper(maxDep, dep)

	traverse_104(root.Left)
	traverse_104(root.Right)

	dep--
}