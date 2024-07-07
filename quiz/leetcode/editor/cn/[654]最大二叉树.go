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
func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	maxIndex := 0
	for i := range nums {
		if nums[i] > nums[maxIndex] {
			maxIndex = i
		}
	}

	root := &TreeNode{Val: nums[maxIndex]}
	root.Left = constructMaximumBinaryTree(nums[0:maxIndex])
	root.Right = constructMaximumBinaryTree(nums[maxIndex+1:])

	return root
}

//leetcode submit region end(Prohibit modification and deletion)
