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
func pseudoPalindromicPaths(root *TreeNode) int {
	pathCount := make([]int, 10)
	var res int
	traverse_1457(root, pathCount, &res)
	return res
}

func traverse_1457(root *TreeNode, pathCount []int, res *int) {
	pathCount[root.Val]++
	defer func() {
		pathCount[root.Val]--
	}()

	if root.Left == nil && root.Right == nil && isPalindromicPath(pathCount) {
		*res++
	}
	if root.Left != nil {
		traverse_1457(root.Left, pathCount, res)
	}
	if root.Right != nil {
		traverse_1457(root.Right, pathCount, res)
	}
}

func isPalindromicPath(pathCount []int) bool {
	var oddCount int
	for _, count := range pathCount {
		if count%2 == 1 {
			oddCount++
			if oddCount > 1 {
				return false
			}
		}
	}
	return true
}
//leetcode submit region end(Prohibit modification and deletion)
