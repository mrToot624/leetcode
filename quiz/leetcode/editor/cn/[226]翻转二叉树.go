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

func invertTree(root *TreeNode) *TreeNode {
	traverseInvert(root)
	return root
}

func traverseInvert(root *TreeNode) {
	if root == nil {
		return
	}

	traverseInvert(root.Left)

	tmp := root.Right
	root.Right = root.Left
	root.Left = tmp

	traverseInvert(root.Left)
}

func splitInvertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	invertedLeft := invertTree(root.Left)
	invertedRight := invertTree(root.Right)

	root.Left = invertedRight
	root.Right = invertedLeft

	return root
}
//leetcode submit region end(Prohibit modification and deletion)
