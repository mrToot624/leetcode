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
func preorderTraversal(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode

	cur := root
	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			res = append(res, cur.Val)
			cur = cur.Left
		}
		cur = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}
	return res
}

func recursivePreorderTraversal(root *TreeNode) []int {
	var res []int
	traverse_144(root, &res)
	return res
}

func traverse_144(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}

	*res = append(*res, root.Val)
	traverse_144(root.Left, res)
	traverse_144(root.Right, res)
}

//leetcode submit region end(Prohibit modification and deletion)

func splitPreorderTraversal(root *TreeNode) []int {
	var preOrderRes []int
	if root == nil {
		return preOrderRes
	}

	preOrderRes = append(preOrderRes, root.Val)
	preOrderRes = append(preOrderRes, splitPreorderTraversal(root.Left)...)
	preOrderRes = append(preOrderRes, splitPreorderTraversal(root.Right)...)

	return preOrderRes
}
