package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 每层相连的二叉树
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}
