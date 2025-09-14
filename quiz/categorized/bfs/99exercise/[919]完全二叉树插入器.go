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
type CBTInserter struct {
	q    []*TreeNode
	root *TreeNode
}

func Constructor_919(root *TreeNode) CBTInserter {
	q := []*TreeNode{root}
	for len(q) > 0 {
		top := q[0]
		if top.Left != nil {
			q = append(q, top.Left)
		} else {
			break
		}
		if top.Right != nil {
			q = append(q, top.Right)
			q = q[1:]
		} else {
			break
		}
	}
	return CBTInserter{q, root}
}

func (this *CBTInserter) Insert(val int) int {
	top := this.q[0]
	newNode := &TreeNode{Val: val}
	if top.Left == nil {
		top.Left = newNode
	} else {
		top.Right = newNode
		this.q = this.q[1:]
	}
	this.q = append(this.q, newNode)
	return top.Val
}

func (this *CBTInserter) Get_root() *TreeNode {
	return this.root
}

/**
 * Your CBTInserter object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Insert(val);
 * param_2 := obj.Get_root();
 */
//leetcode submit region end(Prohibit modification and deletion)
