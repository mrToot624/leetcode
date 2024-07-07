package main

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 1 {
		return head
	}

	reversedHead, isReversed := iterateReverseTopK(head, k)
	if !isReversed {
		return head
	}

	head.Next = reverseKGroup(head.Next, k)
	return reversedHead
}

func iterateReverseKGroup(head *ListNode, k int) *ListNode {
	if k == 1 {
		return head
	}

	dummy := &ListNode{Next: head}
	prevTail, curGroupHead := dummy, head

	for true {
		newHead, isReversed := iterateReverseTopK(curGroupHead, k)
		if !isReversed {
			break
		}

		prevTail.Next = newHead
		prevTail = curGroupHead
		curGroupHead = curGroupHead.Next
	}

	return dummy.Next
}

func iterateReverseTopK(head *ListNode, k int) (*ListNode, bool) {
	nextGroupHead := head
	for i := 0; i < k; i++ {
		if nextGroupHead == nil {
			return head, false
		}
		nextGroupHead = nextGroupHead.Next
	}

	dummy := &ListNode{Next: head}
	cur := dummy.Next
	for i := 0; i < k-1; i++ {
		next := cur.Next
		cur.Next = next.Next
		next.Next = dummy.Next
		dummy.Next = next
	}

	return dummy.Next, true
}
//leetcode submit region end(Prohibit modification and deletion)
