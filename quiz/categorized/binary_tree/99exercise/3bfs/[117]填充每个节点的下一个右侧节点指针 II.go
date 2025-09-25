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

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	q := []*Node{root}
	for len(q) > 0 {
		l := len(q)
		for i := 0; i < l; i++ {
			cur := q[0]
			q = q[1:]
			if i < l-1 {
				cur.Next = q[0]
			}
			if cur.Left != nil {
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			}
		}
	}
	return root
}

//leetcode submit region end(Prohibit modification and deletion)
