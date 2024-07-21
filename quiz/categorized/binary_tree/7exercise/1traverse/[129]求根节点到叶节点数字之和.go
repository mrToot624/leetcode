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
func sumNumbers(root *TreeNode) int {
	var res, pathSum int
	traverse_129(root, &res, pathSum)
	return res
}

func traverse_129(root *TreeNode, res *int, pathSum int) {
	pathSum = pathSum*10 + root.Val
	if root.Left == nil && root.Right == nil {
		*res += pathSum
		return
	}
	if root.Left != nil {
		traverse_129(root.Left, res, pathSum)
	}
	if root.Right != nil {
		traverse_129(root.Right, res, pathSum)
	}
}

//leetcode submit region end(Prohibit modification and deletion)
