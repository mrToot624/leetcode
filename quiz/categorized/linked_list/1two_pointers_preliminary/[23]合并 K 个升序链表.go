package main

import (
	"container/heap"
)

//leetcode submit region begin(Prohibit modification and deletion)
/**
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
 */

type priorityQueue []*ListNode

func (pq priorityQueue) Len() int { return len(pq) }

func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].Val < pq[j].Val
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *priorityQueue) Push(x interface{}) {
	item := x.(*ListNode)
	*pq = append(*pq, item)
}

func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func mergeKLists(lists []*ListNode) *ListNode {
	dummy := &ListNode{}
	p := dummy

	var pq priorityQueue
	for _, head := range lists {
		if head != nil {
			heap.Push(&pq, head)
		}
	}

	for pq.Len() != 0 {
		head := heap.Pop(&pq).(*ListNode)
		if head.Next != nil {
			heap.Push(&pq, head.Next)
			head.Next = nil
		}
		p.Next = head
		p = p.Next
	}

	return dummy.Next
}

//leetcode submit region end(Prohibit modification and deletion)
