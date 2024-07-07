package main

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == right {
		return head
	}

	dummyForWholeList := &ListNode{Next: head}

	dummyForReversed := dummyForWholeList
	for i := 0; i < left-1; i++ {
		dummyForReversed = dummyForReversed.Next
	}

	cur := dummyForReversed.Next
	for i := 0; i < right-left; i++ {
		next := cur.Next
		cur.Next = next.Next
		next.Next = dummyForReversed.Next
		dummyForReversed.Next = next
	}

	return dummyForWholeList.Next
}

var firstDetached *ListNode = nil

func recursiveReverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == 1 {
		return reverseTopK(head, right)
	}
	head.Next = recursiveReverseBetween(head.Next, left-1, right-1)
	return head
}

func reverseTopK(head *ListNode, k int) *ListNode {
	if k == 1 {
		firstDetached = head.Next
		return head
	}

	reversedHead := reverseTopK(head.Next, k-1)

	head.Next.Next = head
	head.Next = firstDetached

	return reversedHead
}

func reverseBetweenUgly(head *ListNode, left int, right int) *ListNode {
	if left == right {
		return head
	}

	var lastBeforeReverse, firstAfterReverse *ListNode
	var firstReversed, lastReversed *ListNode

	dummy := &ListNode{}
	dummy.Next = head

	prev, cur := dummy, head
	for right > 0 {
		next := cur.Next
		if left <= 1 && right >= 1 {
			if left == 1 {
				firstReversed = cur
				lastBeforeReverse = prev
			}

			if right == 1 {
				lastReversed = cur
				firstAfterReverse = next
			}

			if left < 1 {
				cur.Next = prev
			}
		}
		prev = cur
		cur = next
		left--
		right--
	}

	firstReversed.Next = firstAfterReverse
	lastBeforeReverse.Next = lastReversed

	return dummy.Next
}
//leetcode submit region end(Prohibit modification and deletion)
