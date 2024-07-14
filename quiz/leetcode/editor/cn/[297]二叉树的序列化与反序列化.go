package main

import (
	"strconv"
	"strings"
)

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

const (
	sep_297  = ","
	null_297 = "#"
)

//type Codec struct {
//	tree []string
//}

type Codec struct {
	builder *strings.Builder
}

func Constructor_297() Codec {
	return Codec{builder: new(strings.Builder)}
}

// Serializes a tree to a single string.
//func (this *Codec) serialize(root *TreeNode) string {
//	this.serializeBTree(root)
//	return strings.Join(this.tree, sep_297)
//}
//
//func (this *Codec) serializeBTree(root *TreeNode) {
//	if root == nil {
//		this.tree = append(this.tree, null_297)
//		return
//	}
//
//	this.tree = append(this.tree, strconv.Itoa(root.Val))
//
//	this.serializeBTree(root.Left)
//	this.serializeBTree(root.Right)
//}

func (this *Codec) serialize(root *TreeNode) string {
	this.serializeBTree(root)
	return strings.TrimSuffix(this.builder.String(), sep_297)
}

func (this *Codec) serializeBTree(root *TreeNode) {
	if root == nil {
		this.builder.WriteString(null_297)
		this.builder.WriteString(sep_297)
		return
	}

	this.builder.WriteString(strconv.Itoa(root.Val))
	this.builder.WriteString(sep_297)

	this.serializeBTree(root.Left)
	this.serializeBTree(root.Right)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	tree := strings.Split(data, sep_297)
	return deserializeBTree(&tree)
}

func deserializeBTree(tree *[]string) *TreeNode {
	rootStr := (*tree)[0]
	*tree = (*tree)[1:]

	if rootStr == null_297 {
		return nil
	}

	rootVal, _ := strconv.Atoi(rootStr)
	root := &TreeNode{Val: rootVal}

	root.Left = deserializeBTree(tree)
	root.Right = deserializeBTree(tree)

	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
//leetcode submit region end(Prohibit modification and deletion)
