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
func postorderTraversal(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode

	cur, prev := root, root
	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}

		r := stack[len(stack)-1]
		if r.Right == nil || prev == r.Right {
			stack = stack[:len(stack)-1]
			res = append(res, r.Val)
			prev = r
		} else {
			cur = r.Right
		}
	}
	return res
}

func recursivePostorderTraversal(root *TreeNode) []int {
	var res []int
	traverse_145(root, &res)
	return res
}

func traverse_145(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}

	traverse_145(root.Left, res)
	traverse_145(root.Right, res)
	*res = append(*res, root.Val)
}

//leetcode submit region end(Prohibit modification and deletion)
