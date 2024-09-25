package _preliminary

// leetcode submit region begin(Prohibit modification and deletion)
type queue struct {
	l []int
}

func (q *queue) push(x int) {
	q.l = append(q.l, x)
}

func (q *queue) pop() int {
	res := q.l[0]
	q.l = q.l[1:]
	return res
}

func (q *queue) len() int {
	return len(q.l)
}

func (q *queue) peek() int {
	return q.l[0]
}

type MyStack struct {
	q queue
}

func Constructor_225() MyStack {
	return MyStack{}
}

func (this *MyStack) Push(x int) {
	this.q.push(x)
	if this.q.len() > 1 {
		for i := 0; i < this.q.len()-1; i++ {
			this.q.push(this.q.pop())
		}
	}
}

func (this *MyStack) Pop() int {
	return this.q.pop()
}

func (this *MyStack) Top() int {
	return this.q.peek()
}

func (this *MyStack) Empty() bool {
	return this.q.len() == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
//leetcode submit region end(Prohibit modification and deletion)
