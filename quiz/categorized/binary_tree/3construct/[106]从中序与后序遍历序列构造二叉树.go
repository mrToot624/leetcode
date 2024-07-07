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
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}

	if len(postorder) == 1 {
		return &TreeNode{Val: postorder[0]}
	}

	rootIndex := 0
	for i, num := range inorder {
		if num == postorder[len(postorder)-1] {
			rootIndex = i
			break
		}
	}

	root := &TreeNode{Val: postorder[len(postorder)-1]}
	root.Left = buildTree(inorder[:rootIndex], postorder[:rootIndex])
	root.Right = buildTree(inorder[rootIndex+1:], postorder[rootIndex:len(postorder)-1])

	return root
}
//leetcode submit region end(Prohibit modification and deletion)
