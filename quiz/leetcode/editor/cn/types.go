package main

// 链表
type ListNode struct {
	Val  int
	Next *ListNode
}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 完美二叉树
//type Node struct {
//	Val   int
//	Left  *Node
//	Right *Node
//	Next  *Node
//}

// N叉树
type Node struct {
	Val      int
	Children []*Node
}

// This is the interface that allows for creating nested lists.
// You should not implement it, or speculate about its implementation
type NestedInteger struct {
}

// Return true if this NestedInteger holds a single integer, rather than a nested list.
func (this NestedInteger) IsInteger() bool { return false }

// Return the single integer that this NestedInteger holds, if it holds a single integer
// The result is undefined if this NestedInteger holds a nested list
// So before calling this method, you should have a check
func (this NestedInteger) GetInteger() int { return 0 }

// Set this NestedInteger to hold a single integer.
func (n *NestedInteger) SetInteger(value int) {}

// Set this NestedInteger to hold a nested list and adds a nested integer to it.
func (this *NestedInteger) Add(elem NestedInteger) {}

// Return the nested list that this NestedInteger holds, if it holds a nested list
// The list length is zero if this NestedInteger holds a single integer
// You can access NestedInteger's List element directly if you want to modify it
func (this NestedInteger) GetList() []*NestedInteger { return nil }

// 差分数组
type difference struct {
	diff []int
}

func NewDifference(nums []int) *difference {
	diff := make([]int, len(nums))

	diff[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		diff[i] = nums[i] - nums[i-1]
	}
	return &difference{diff}
}

func (d *difference) increment(start, end, delta int) {
	d.diff[start] += delta
	if end+1 < len(d.diff) {
		d.diff[end+1] -= delta
	}
}

func (d *difference) result() []int {
	res := make([]int, len(d.diff))
	res[0] = d.diff[0]
	for i := 1; i < len(res); i++ {
		res[i] = res[i-1] + d.diff[i]
	}
	return res
}

// Union Find
type UnionFind struct {
	parent []int
	count  int
}

func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	return &UnionFind{parent, n}
}

func (uf *UnionFind) find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *UnionFind) Union(p, q int) {
	pRoot := uf.find(p)
	qRoot := uf.find(q)
	if pRoot != qRoot {
		uf.parent[pRoot] = qRoot
	}
	uf.count--
}

func (uf *UnionFind) Connected(p, q int) bool {
	return uf.find(p) == uf.find(q)
}

func (uf *UnionFind) Count() int {
	return uf.count
}
