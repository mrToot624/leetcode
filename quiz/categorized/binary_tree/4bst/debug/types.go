package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 完美二叉树
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}
