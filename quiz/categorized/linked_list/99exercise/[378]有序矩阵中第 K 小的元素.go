package main

import (
	"container/heap"
)

//leetcode submit region begin(Prohibit modification and deletion)

type pq [][3]int

func (h pq) Less(i, j int) bool {
	return h[i][0] < h[j][0]
}

func (h pq) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *pq) Push(x any) {
	*h = append(*h, x.([3]int))
}

func (h *pq) Pop() any {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

func (h pq) Len() int {
	return len(h)
}

func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix)
	q := &pq{}
	for i := 0; i < n; i++ {
		heap.Push(q, [3]int{matrix[i][0], i, 0})
	}

	var res int
	for i := 0; i < k; i++ {
		top := heap.Pop(q).([3]int)
		res = top[0]
		row, column := top[1], top[2]
		if column < n-1 {
			heap.Push(q, [3]int{matrix[row][column+1], row, column + 1})
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
