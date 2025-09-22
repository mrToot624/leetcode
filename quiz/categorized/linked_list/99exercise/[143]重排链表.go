package main

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	var l []*ListNode
	cur := head
	for cur != nil {
		l = append(l, cur)
		cur = cur.Next
	}

	left, right := 0, len(l)-1
	for left+1 < right {
		l[right-1].Next = nil
		l[right].Next = l[left].Next
		l[left].Next = l[right]
		right--
		left++
	}
}

//leetcode submit region end(Prohibit modification and deletion)
