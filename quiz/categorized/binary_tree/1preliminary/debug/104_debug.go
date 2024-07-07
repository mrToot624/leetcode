package main

import "fmt"

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var depth, res int

func maxDepth(root *TreeNode) int {
	if root == nil {
		return res
	}

	traverse(root)

	return res
}

func traverse(root *TreeNode) {
	if root == nil {
		return
	}

	depth++
	res = getDeeper(res, depth)

	traverse(root.Left)
	traverse(root.Right)

	depth--
}

func getDeeper(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	root := &TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}

	right := &TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}
	root.Right = right
	fmt.Println(maxDepth(root))
}

//leetcode submit region end(Prohibit modification and deletion)
