package main

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a NTree.
 * type NTree struct {
 *     Val int
 *     Children []*NTree
 * }
 */

func postorder(root *NTree) []int {
	var res []int
	var traverse func(root *NTree)
	traverse = func(root *NTree) {
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
