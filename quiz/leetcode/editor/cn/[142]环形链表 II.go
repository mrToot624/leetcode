package main

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	fast, slow := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			break
		}
	}

	// fast能走到终点，说明无环
	if fast == nil || fast.Next == nil {
		return nil
	}

	slow = head
	for fast != slow {
		slow = slow.Next
		fast = fast.Next
	}
	return fast
}
//leetcode submit region end(Prohibit modification and deletion)
