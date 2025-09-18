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
func rightSideView(root *TreeNode) []int {
	var res []int
	traverse_199(root, &res, 1)
	return res
}

func traverse_199(root *TreeNode, res *[]int, depth int) {
	if root == nil {
		return
	}

	if len(*res) < depth {
		*res = append(*res, root.Val)
	}
	traverse_199(root.Right, res, depth+1)
	traverse_199(root.Left, res, depth+1)
}

//leetcode submit region end(Prohibit modification and deletion)
