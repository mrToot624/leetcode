package main

import "math/rand"

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type Solution_382 struct {
	head *ListNode
}

func Constructor_382(head *ListNode) Solution_382 {
	return Solution_382{head}
}

func (this *Solution_382) GetRandom() int {
	cur := this.head
	i := 1
	res := this.head.Val
	for cur != nil {
		if rand.Intn(i) == 0 {
			res = cur.Val
		}
		cur = cur.Next
		i++
	}
	return res
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */
//leetcode submit region end(Prohibit modification and deletion)
