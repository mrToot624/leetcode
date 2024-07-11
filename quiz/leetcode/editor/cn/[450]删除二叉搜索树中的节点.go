package main

import "math"

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deleteNode(root *TreeNode, key int) *TreeNode {
	dummy := &TreeNode{Val: math.MaxInt, Left: root}

	prev, cur := dummy, root
	for cur != nil {
		if cur.Val == key {
			if cur.Right == nil {
				connectSubstitute(prev, cur.Left, key)
			} else if cur.Left == nil {
				connectSubstitute(prev, cur.Right, key)
			} else {
				substitute := findMinFromRightChild(cur)
				connectSubstitute(prev, substitute, key)
				substitute.Left = cur.Left
				substitute.Right = cur.Right
			}
			break
		} else {
			prev = cur
			if cur.Val > key {
				cur = cur.Left
			} else {
				cur = cur.Right
			}
		}
	}

	return dummy.Left
}

func connectSubstitute(prev, substitute *TreeNode, key int) {
	if prev.Val > key {
		prev.Left = substitute
	} else {
		prev.Right = substitute
	}
}

func findMinFromRightChild(root *TreeNode) *TreeNode {
	prev, cur := root, root.Right
	for cur.Left != nil {
		prev = cur
		cur = cur.Left
	}

	if prev == root {
		prev.Right = cur.Right
	} else {
		prev.Left = cur.Right
	}
	return cur
}

//leetcode submit region end(Prohibit modification and deletion)
