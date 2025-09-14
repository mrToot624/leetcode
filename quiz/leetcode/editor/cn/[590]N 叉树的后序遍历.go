package main

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func postorder(root *Node) []int {
    var res []int
	var traverse func(root *Node)
	traverse = func(root *Node) {
		if root == nil {
			return
		}

		for _, child := range root.Children {
			traverse(child)
		}
		res = append(res, root.Val)
	}
	traverse(root)
	return res
}
//leetcode submit region end(Prohibit modification and deletion)
