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
func kthSmallest(root *TreeNode, k int) int {
	var count, res int
	traverse_230(root, k, &count, &res)
	return res
}

func traverse_230(root *TreeNode, k int, count, res *int) {
	if root == nil {
		return
	}

	traverse_230(root.Left, k, count, res)

	*count++
	if *count >= k {
		if *count == k {
			*res = root.Val
		}
		return
	}

	traverse_230(root.Right, k, count, res)
}

//leetcode submit region end(Prohibit modification and deletion)
