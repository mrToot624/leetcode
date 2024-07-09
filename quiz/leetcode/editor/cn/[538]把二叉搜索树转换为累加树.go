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
func convertBST(root *TreeNode) *TreeNode {
	var sum int
	traverse_538(root, &sum)
	return root
}

func traverse_538(root *TreeNode, sum *int) {
	if root == nil {
		return
	}

	traverse_538(root.Right, sum)

	*sum += root.Val
	root.Val = *sum

	traverse_538(root.Left, sum)
}

//leetcode submit region end(Prohibit modification and deletion)
