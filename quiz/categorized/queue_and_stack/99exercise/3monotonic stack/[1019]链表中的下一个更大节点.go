package main

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func nextLargerNodes(head *ListNode) []int {
	cur := head
	var nums, mq []int
	for cur != nil {
		nums = append(nums, cur.Val)
		cur = cur.Next
	}

	res := make([]int, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		for len(mq) > 0 && mq[len(mq)-1] <= nums[i] {
			mq = mq[:len(mq)-1]
		}
		if len(mq) == 0 {
			res[i] = 0
		} else {
			res[i] = mq[len(mq)-1]
		}
		mq = append(mq, nums[i])
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
