package main

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (this NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (this NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (this *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (this NestedInteger) GetList() []*NestedInteger {}
 */

type NestedIterator struct {
	stack []*NestedInteger
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	stack := append([]*NestedInteger{}, nestedList...)
	return &NestedIterator{stack}
}

func (this *NestedIterator) Next() int {
	res := this.stack[0].GetInteger()
	this.stack = this.stack[1:]
	return res
}

func (this *NestedIterator) HasNext() bool {
	for len(this.stack) > 0 {
		if this.stack[0].IsInteger() {
			return true
		}
		firstChildren := this.stack[0].GetList()
		this.stack = append(firstChildren, this.stack[1:]...)
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)
