package main

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	dummy := &ListNode{Next: slow}
	for slow.Next != nil {
		next := slow.Next
		slow.Next = next.Next
		next.Next = dummy.Next
		dummy.Next = next
	}

	p1, p2 := head, dummy.Next
	for p2 != nil {
		if p1.Val != p2.Val {
			return false
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	return true
}
//leetcode submit region end(Prohibit modification and deletion)
