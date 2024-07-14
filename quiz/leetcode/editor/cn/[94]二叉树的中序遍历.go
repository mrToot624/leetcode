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
func inorderTraversal(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode

	cur := root
	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}

		r := stack[len(stack)-1]
		res = append(res, r.Val)
		stack = stack[:len(stack)-1]
		cur = r.Right
	}
	return res
}

func recursiveInorderTraversal(root *TreeNode) []int {
	var res []int
	traverse_94(root, &res)
	return res
}

func traverse_94(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}

	traverse_94(root.Left, res)
	*res = append(*res, root.Val)
	traverse_94(root.Right, res)
}

//leetcode submit region end(Prohibit modification and deletion)
