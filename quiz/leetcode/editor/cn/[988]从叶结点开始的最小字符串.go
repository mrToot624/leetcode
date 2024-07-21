package main

import "strings"

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func smallestFromLeaf(root *TreeNode) string {
	res := string('z' + 1)
	var track []rune
	traverse_988(root, &track, &res)
	return res
}

func traverse_988(root *TreeNode, track *[]rune, res *string) {
	*track = append(*track, 'a'+rune(root.Val))
	defer func() {
		*track = (*track)[:len(*track)-1]
	}()

	if root.Left == nil && root.Right == nil {
		str := constructPathString(track)
		if str < *res {
			*res = str
		}
		return
	}

	if root.Left != nil {
		traverse_988(root.Left, track, res)
	}
	if root.Right != nil {
		traverse_988(root.Right, track, res)
	}
}

func constructPathString(track *[]rune) string {
	var res strings.Builder
	for i := len(*track) - 1; i >= 0; i-- {
		res.WriteRune((*track)[i])
	}
	return res.String()
}

//leetcode submit region end(Prohibit modification and deletion)
