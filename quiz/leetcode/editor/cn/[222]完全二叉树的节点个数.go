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
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var depth int
	cur := root.Left
	for cur != nil {
		depth++
		cur = cur.Left
	}

	left, right := 1<<depth, 1<<(depth+1)-1
	for left <= right {
		mid := left + (right-left)/2
		if leafExists(root, mid) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return right
}

func leafExists(root *TreeNode, num int) bool {
	var op []rune
	for num != 1 {
		if num%2 == 0 {
			op = append(op, 'l')
		} else {
			op = append(op, 'r')
		}
		num >>= 1
	}

	cur := root
	for i := len(op) - 1; i >= 0; i-- {
		if op[i] == 'l' {
			cur = cur.Left
		} else {
			cur = cur.Right
		}
	}
	return cur != nil
}

//leetcode submit region end(Prohibit modification and deletion)
