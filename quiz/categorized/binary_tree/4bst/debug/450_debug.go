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
		prev.Left = cur.Left
	}
	return cur
}

func main() {
	two := &TreeNode{Val: 2}
	four := &TreeNode{Val: 4}
	seven := &TreeNode{Val: 7}

	three := &TreeNode{Val: 3, Left: two, Right: four}
	six := &TreeNode{Val: 6, Right: seven}

	root := &TreeNode{Val: 5, Left: three, Right: six}

	deleteNode(root, 3)
}

//leetcode submit region end(Prohibit modification and deletion)
