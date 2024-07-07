package main

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head

	target := findNthFromEnd(dummy, n+1)
	target.Next = target.Next.Next

	return dummy.Next
}

func findNthFromEnd(head *ListNode, n int) *ListNode {
	fastP, slowP := head, head
	for i := 0; i < n; i++ {
		fastP = fastP.Next
	}

	for fastP != nil {
		fastP = fastP.Next
		slowP = slowP.Next
	}
	return slowP
}

//leetcode submit region end(Prohibit modification and deletion)
