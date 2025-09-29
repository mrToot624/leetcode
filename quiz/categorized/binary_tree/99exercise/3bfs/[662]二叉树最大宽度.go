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
func widthOfBinaryTree(root *TreeNode) int {
	type indexedTNode struct {
		node  *TreeNode
		index int
	}
	var maxWidth int
	q := []indexedTNode{{node: root, index: 1}}
	for len(q) > 0 {
		l := len(q)
		if l >= 1 {
			maxWidth = max(maxWidth, q[l-1].index-q[0].index+1)
		}
		for i := 0; i < l; i++ {
			cur := q[0]
			q = q[1:]
			index := cur.index
			if cur.node.Left != nil {
				q = append(q, indexedTNode{node: cur.node.Left, index: 2 * index})
			}
			if cur.node.Right != nil {
				q = append(q, indexedTNode{node: cur.node.Right, index: 2*index + 1})
			}
		}
	}
	return maxWidth
}

//leetcode submit region end(Prohibit modification and deletion)
