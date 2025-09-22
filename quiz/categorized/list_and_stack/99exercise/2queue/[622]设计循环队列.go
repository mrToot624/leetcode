package main

// leetcode submit region begin(Prohibit modification and deletion)
type MyCircularQueue struct {
	cq         []int
	head, tail int
	size       int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		cq: make([]int, k),
	}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	this.cq[this.tail] = value
	this.tail = (this.tail + 1) % len(this.cq)
	this.size++
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	this.head = (this.head + 1) % len(this.cq)
	this.size--
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.cq[this.head]
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.cq[(this.tail-1+len(this.cq))%len(this.cq)]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.size == 0
}

func (this *MyCircularQueue) IsFull() bool {
	return this.size == len(this.cq)
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
//leetcode submit region end(Prohibit modification and deletion)
