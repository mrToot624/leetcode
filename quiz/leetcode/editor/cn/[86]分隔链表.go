package main

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition_86(head *ListNode, x int) *ListNode {
	smallDummy, bigDummy := &ListNode{}, &ListNode{}
	s, b, cur := smallDummy, bigDummy, head

	for cur != nil {
		if cur.Val < x {
			s.Next = cur
			s = s.Next
		} else {
			b.Next = cur
			b = b.Next
		}
		cur = cur.Next
	}

	s.Next = bigDummy.Next
	b.Next = nil

	return smallDummy.Next
}

//leetcode submit region end(Prohibit modification and deletion)
