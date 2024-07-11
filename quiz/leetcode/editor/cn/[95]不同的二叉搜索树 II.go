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
func generateTrees(n int) []*TreeNode {
	dp := make([][][]*TreeNode, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = make([][]*TreeNode, n+1)
		dp[i][i] = []*TreeNode{{Val: i}}
	}

	for L := 1; L < n; L++ {
		for j := 1; j+L <= n; j++ {
			var trees []*TreeNode
			for r := j; r <= j+L; r++ {
				for _, left := range dpWithValidBoundary(dp, j, r-1) {
					for _, right := range dpWithValidBoundary(dp, r+1, j+L) {
						trees = append(trees, &TreeNode{r, left, right})
					}
				}
			}
			dp[j][j+L] = trees
		}
	}
	return dp[1][n]
}

func dpWithValidBoundary(dp [][][]*TreeNode, start, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}
	return dp[start][end]
}

//leetcode submit region end(Prohibit modification and deletion)
