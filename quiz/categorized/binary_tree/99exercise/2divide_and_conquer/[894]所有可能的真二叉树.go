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
func allPossibleFBT(n int) []*TreeNode {
	dp := make([][]*TreeNode, n+1)
	dp[1] = []*TreeNode{{}}
	for i := 3; i <= n; i += 2 {
		var trees []*TreeNode
		for j := 1; j <= i-1; j += 2 {
			for _, left := range dp[j] {
				for _, right := range dp[i-1-j] {
					trees = append(trees, &TreeNode{Left: left, Right: right})
				}
			}
		}
		dp[i] = trees
	}
	return dp[n]
}

//leetcode submit region end(Prohibit modification and deletion)
