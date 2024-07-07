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
func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	if len(preorder) == 1 {
		return &TreeNode{Val: preorder[0]}
	}

	presumedLeftChildIndex := 0
	for i, num := range postorder {
		if num == preorder[1] {
			presumedLeftChildIndex = i
			break
		}
	}

	root := &TreeNode{Val: preorder[0]}
	root.Left = constructFromPrePost(preorder[1:presumedLeftChildIndex+2], postorder[:presumedLeftChildIndex+1])
	root.Right = constructFromPrePost(preorder[presumedLeftChildIndex+2:], postorder[presumedLeftChildIndex+1:len(postorder)-1])

	return root
}

//leetcode submit region end(Prohibit modification and deletion)
