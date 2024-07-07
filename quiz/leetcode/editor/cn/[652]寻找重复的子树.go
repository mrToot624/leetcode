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
const (
	sep_652  = ","
	null_652 = "#"
)

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	var res []*TreeNode
	treeSet := make(map[string]bool)

	traverseAndSerialize(root, treeSet, &res)

	return res
}

func traverseAndSerialize(root *TreeNode, treeSet map[string]bool, res *[]*TreeNode) string {
	var builder strings.Builder
	if root == nil {
		builder.WriteString(null_652)
		builder.WriteString(sep_652)
		return builder.String()
	}

	builder.WriteString(strconv.Itoa(root.Val))
	builder.WriteString(sep_652)

	builder.WriteString(traverseAndSerialize(root.Left, treeSet, res))
	builder.WriteString(traverseAndSerialize(root.Right, treeSet, res))

	if added, ok := treeSet[builder.String()]; ok {
		if !added {
			*res = append(*res, root)
			treeSet[builder.String()] = true
		}
	} else {
		treeSet[builder.String()] = false
	}

	return builder.String()
}

//leetcode submit region end(Prohibit modification and deletion)
