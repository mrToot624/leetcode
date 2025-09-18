package main

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates_83(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	slow, fast := head, head.Next
	for fast != nil {
		if slow.Val == fast.Val {
			slow.Next = fast.Next
			fast = fast.Next
		} else {
			slow = slow.Next
			fast = fast.Next
		}
	}

	return head
}

//leetcode submit region end(Prohibit modification and deletion)
