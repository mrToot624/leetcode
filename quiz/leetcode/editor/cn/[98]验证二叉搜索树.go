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
func isValidBST(root *TreeNode) bool {
	return validateBSTByBoundary(root, nil, nil)
}

func validateBSTByBoundary(root *TreeNode, min, max *int) bool {
	if root == nil {
		return true
	}

	if min != nil && root.Val <= *min {
		return false
	}
	if max != nil && root.Val >= *max {
		return false
	}

	return validateBSTByBoundary(root.Left, min, &root.Val) && validateBSTByBoundary(root.Right, &root.Val, max)
}

//leetcode submit region end(Prohibit modification and deletion)
