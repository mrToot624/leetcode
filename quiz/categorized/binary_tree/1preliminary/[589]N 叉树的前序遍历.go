package main

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func preorder(root *Node) []int {
    var res []int
	var traverse func(root *Node)
	traverse = func(node *Node) {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		for _, child := range node.Children {
			traverse(child)
		}
	}
	traverse(root)
	return res
}
//leetcode submit region end(Prohibit modification and deletion)
