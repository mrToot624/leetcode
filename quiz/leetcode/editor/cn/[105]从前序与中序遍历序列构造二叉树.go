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

// remove the suffix when resubmitting
func buildTree_105(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	if len(preorder) == 1 {
		return &TreeNode{Val: preorder[0]}
	}

	rootIndex := 0
	for i, num := range inorder {
		if num == preorder[0] {
			rootIndex = i
			break
		}
	}

	root := &TreeNode{Val: preorder[0]}
	root.Left = buildTree_105(preorder[1:rootIndex+1], inorder[:rootIndex])
	root.Right = buildTree_105(preorder[rootIndex+1:], inorder[rootIndex+1:])

	return root
}

//leetcode submit region end(Prohibit modification and deletion)
