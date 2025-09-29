package main

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	p1, p2 := l1, l2
	dummy := &ListNode{}
	cur := dummy
	carry := 0
	for p1 != nil || p2 != nil {
		sum := carry
		if p1 != nil {
			sum += p1.Val
			p1 = p1.Next
		}
		if p2 != nil {
			sum += p2.Val
			p2 = p2.Next
		}
		if sum >= 10 {
			sum -= 10
			carry = 1
		} else {
			carry = 0
		}
		cur.Next = &ListNode{Val: sum}
		cur = cur.Next
	}
	if carry == 1 {
		cur.Next = &ListNode{Val: carry}
	}
	return dummy.Next
}

//leetcode submit region end(Prohibit modification and deletion)
