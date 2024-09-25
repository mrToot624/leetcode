package main

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}

	l := 1
	cur := head
	for cur.Next != nil {
		l++
		cur = cur.Next
	}

	shift := k%l
	if shift == 0 {
		return head
	}

	cur.Next = head
	for i := 0; i < l-shift; i++ {
		cur = cur.Next
	}
	newHead := cur.Next
	cur.Next = nil
	return newHead
}

//leetcode submit region end(Prohibit modification and deletion)
