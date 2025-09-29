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
func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	toDelete := make(map[int]bool)
	for _, val := range to_delete {
		toDelete[val] = true
	}
	var res []*TreeNode

	var dfs func(node *TreeNode, parentDeleted bool)
	dfs = func(node *TreeNode, parentDeleted bool) {
		if node == nil {
			return
		}
		if parentDeleted && !toDelete[node.Val] {
			res = append(res, node)
		}
		dfs(node.Left, toDelete[node.Val])
		dfs(node.Right, toDelete[node.Val])
		if node.Left != nil && toDelete[node.Left.Val] {
			node.Left = nil
		}
		if node.Right != nil && toDelete[node.Right.Val] {
			node.Right = nil
		}
	}
	dfs(root, true)
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
