package main

// leetcode submit region begin(Prohibit modification and deletion)
type DoubleListNode struct {
	Prev, Next *DoubleListNode
	Val        int
}
type FrontMiddleBackQueue struct {
	head, tail, middle *DoubleListNode
	size               int
}

func Constructor() FrontMiddleBackQueue {
	return FrontMiddleBackQueue{}
}

func (this *FrontMiddleBackQueue) isEmpty() bool {
	return this.size == 0
}

func (this *FrontMiddleBackQueue) init(val int) {
	this.head = &DoubleListNode{Val: val}
	this.tail, this.middle = this.head, this.head
	this.size++
}

func (this *FrontMiddleBackQueue) increaseMidBySize() {
	if this.size%2 == 1 {
		this.middle = this.middle.Next
	}
}

func (this *FrontMiddleBackQueue) decreaseMidBySize() {
	if this.size%2 == 0 {
		this.middle = this.middle.Prev
	}
}

func (this *FrontMiddleBackQueue) clear() {
	this.head, this.tail, this.middle = nil, nil, nil
	this.size = 0
}

func (this *FrontMiddleBackQueue) PushFront(val int) {
	if this.isEmpty() {
		this.init(val)
		return
	}
	newHead := &DoubleListNode{Val: val, Next: this.head}
	this.head.Prev = newHead
	this.head = newHead
	this.size++
	this.decreaseMidBySize()
}

func (this *FrontMiddleBackQueue) PushMiddle(val int) {
	if this.middle == this.tail {
		this.PushFront(val)
		return
	}
	this.size++
	this.decreaseMidBySize()
	newNode := &DoubleListNode{Val: val, Prev: this.middle, Next: this.middle.Next}
	this.middle.Next.Prev = newNode
	this.middle.Next = newNode
	this.middle = newNode
}

func (this *FrontMiddleBackQueue) PushBack(val int) {
	if this.isEmpty() {
		this.init(val)
		return
	}
	newTail := &DoubleListNode{Val: val, Prev: this.tail}
	this.tail.Next = newTail
	this.tail = newTail
	this.size++
	this.increaseMidBySize()
}

func (this *FrontMiddleBackQueue) PopFront() int {
	if this.isEmpty() {
		return -1
	}
	val := this.head.Val
	this.size--
	if this.isEmpty() {
		this.clear()
	} else {
		this.head.Next.Prev = nil
		this.head = this.head.Next
		this.increaseMidBySize()
	}
	return val
}

func (this *FrontMiddleBackQueue) PopMiddle() int {
	if this.isEmpty() {
		return -1
	}
	val := this.middle.Val
	if this.middle == this.head {
		return this.PopFront()
	}
	this.size--
	if this.isEmpty() {
		this.clear()
	} else {
		this.middle.Prev.Next = this.middle.Next
		this.middle.Next.Prev = this.middle.Prev
		if this.size%2 == 1 {
			this.middle = this.middle.Next
		} else {
			this.middle = this.middle.Prev
		}
	}
	return val
}

func (this *FrontMiddleBackQueue) PopBack() int {
	if this.isEmpty() {
		return -1
	}
	val := this.tail.Val
	this.size--
	if this.isEmpty() {
		this.clear()
	} else {
		this.tail.Prev.Next = nil
		this.tail = this.tail.Prev
		this.decreaseMidBySize()
	}
	return val
}

/**
 * Your FrontMiddleBackQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PushFront(val);
 * obj.PushMiddle(val);
 * obj.PushBack(val);
 * param_4 := obj.PopFront();
 * param_5 := obj.PopMiddle();
 * param_6 := obj.PopBack();
 */
//leetcode submit region end(Prohibit modification and deletion)
