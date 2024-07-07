package main

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	reversedHead := reverseList(head.Next)

	head.Next.Next = head
	head.Next = nil

	return reversedHead
}

func iterateReverseList(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head

	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}

func iterateReverseListWithDummy(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	dummy := &ListNode{Next: head}
	cur := dummy.Next

	for cur.Next != nil {
		next := cur.Next
		cur.Next = next.Next
		next.Next = dummy.Next
		dummy.Next = next
	}
	return dummy.Next
}
//leetcode submit region end(Prohibit modification and deletion)
