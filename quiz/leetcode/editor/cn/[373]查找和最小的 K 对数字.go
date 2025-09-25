package main

import "container/heap"

// leetcode submit region begin(Prohibit modification and deletion)
type pq_373 struct {
	q [][]int
}

func (pq pq_373) Len() int {
	return len(pq.q)
}

func (pq pq_373) Less(i, j int) bool {
	return pq.q[i][0]+pq.q[i][1] < pq.q[j][0]+pq.q[j][1]
}

func (pq pq_373) Swap(i, j int) {
	pq.q[i], pq.q[j] = pq.q[j], pq.q[i]
}

func (pq *pq_373) Push(x any) {
	pq.q = append(pq.q, x.([]int))
}

func (pq *pq_373) Pop() any {
	x := pq.q[len(pq.q)-1]
	pq.q = pq.q[:len(pq.q)-1]
	return x
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	var initQ [][]int
	for i := 0; i < min(len(nums1), k); i++ {
		initQ = append(initQ, []int{nums1[i], nums2[0], 0})
	}
	q := &pq_373{initQ}
	heap.Init(q)

	var res [][]int
	for i := k; i > 0; i-- {
		top := heap.Pop(q).([]int)
		res = append(res, []int{top[0], top[1]})
		nextNum2I := top[2] + 1
		if nextNum2I < len(nums2) {
			heap.Push(q, []int{top[0], nums2[nextNum2I], nextNum2I})
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
