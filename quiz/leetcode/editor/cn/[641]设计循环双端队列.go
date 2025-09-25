package main

// leetcode submit region begin(Prohibit modification and deletion)
type MyCircularDeque struct {
	head, tail int
	deque      []int
}

func Constructor_641(k int) MyCircularDeque {
	return MyCircularDeque{deque: make([]int, k+1)}
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}
	this.head = (this.head - 1 + len(this.deque)) % len(this.deque)
	this.deque[this.head] = value
	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}
	this.deque[this.tail] = value
	this.tail = (this.tail + 1) % len(this.deque)
	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	this.head = (this.head + 1) % len(this.deque)
	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	this.tail = (this.tail - 1 + len(this.deque)) % len(this.deque)
	return true
}

func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}
	return this.deque[this.head]
}

func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.deque[(this.tail-1+len(this.deque))%len(this.deque)]
}

func (this *MyCircularDeque) IsEmpty() bool {
	return this.head == this.tail
}

func (this *MyCircularDeque) IsFull() bool {
	return (this.tail+1)%len(this.deque) == this.head
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
//leetcode submit region end(Prohibit modification and deletion)
