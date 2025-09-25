package main

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a NTree.
 * type NTree struct {
 *     Val int
 *     Left *NTree
 *     Right *NTree
 *     Next *NTree
 * }
 */

func connect_116(root *Node) *Node {
	if root == nil {
		return root
	}

	traverseConnect(root.Left, root.Right)
	return root
}

func traverseConnect(left, right *Node) {
	if left == nil || right == nil {
		return
	}

	left.Next = right

	traverseConnect(left.Left, left.Right)
	traverseConnect(left.Right, right.Left)
	traverseConnect(right.Left, right.Right)
}

func bfsConnect(root *Node) *Node {
	if root == nil {
		return root
	}

	var q []*Node
	q = append(q, root)

	for len(q) > 0 {
		currentLevelLen := len(q)
		for i := 0; i < currentLevelLen; i++ {
			current := q[0]
			q = q[1:]

			if i < currentLevelLen-1 {
				current.Next = q[0]
			}

			if current.Left != nil {
				q = append(q, current.Left)
			}
			if current.Right != nil {
				q = append(q, current.Right)
			}
		}
	}
	return root
}

//leetcode submit region end(Prohibit modification and deletion)
