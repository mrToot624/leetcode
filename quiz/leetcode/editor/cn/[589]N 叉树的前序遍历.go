package main

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a NTree.
 * type NTree struct {
 *     Val int
 *     Children []*NTree
 * }
 */

func preorder(root *NTree) []int {
	var res []int
	var traverse func(root *NTree)
	traverse = func(node *NTree) {
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
