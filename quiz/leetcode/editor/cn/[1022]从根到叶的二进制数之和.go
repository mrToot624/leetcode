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
func sumRootToLeaf(root *TreeNode) int {
	var res int
	traverse_1022(root, 0, &res)
	return res
}

func traverse_1022(root *TreeNode, pathSum int, res *int) {
	pathSum = pathSum<<1 | root.Val //pathSum = 2*pathSum + root.Val
	if root.Left == nil && root.Right == nil {
		*res += pathSum
	}
	if root.Left != nil {
		traverse_1022(root.Left, pathSum, res)
	}
	if root.Right != nil {
		traverse_1022(root.Right, pathSum, res)
	}
}
//leetcode submit region end(Prohibit modification and deletion)
