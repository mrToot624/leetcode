package main

import (
	"strconv"
	"strings"
)

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths(root *TreeNode) []string {
	var res, track []string
	traverse_257(root, &res, track)
	return res
}

func traverse_257(root *TreeNode, res *[]string, track []string) {
	track = append(track, strconv.Itoa(root.Val))

	if root.Left == nil && root.Right == nil {
		*res = append(*res, strings.Join(track, "->"))
		return
	}

	if root.Left != nil {
		traverse_257(root.Left, res, track)
	}
	if root.Right != nil {
		traverse_257(root.Right, res, track)
	}
	track = track[:len(track)-1]
}
//leetcode submit region end(Prohibit modification and deletion)
