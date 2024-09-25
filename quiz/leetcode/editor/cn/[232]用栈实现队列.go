package main

// leetcode submit region begin(Prohibit modification and deletion)
type stack struct {
	l []int
}

func (s *stack) len() int {
	return len(s.l)
}

func (s *stack) push(x int) {
	s.l = append(s.l, x)
}

func (s *stack) pop() int {
	res := s.l[s.len()-1]
	s.l = s.l[0 : s.len()-1]
	return res
}

func (s *stack) peek() int {
	return s.l[s.len()-1]
}

type MyQueue struct {
	s1 stack
	s2 stack
}

func Constructor_232() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) shift() {
	if this.s2.len() == 0 {
		for this.s1.len() > 0 {
			this.s2.push(this.s1.pop())
		}
	}
}

func (this *MyQueue) Push(x int) {
	this.s1.push(x)
}

func (this *MyQueue) Pop() int {
	this.shift()
	return this.s2.pop()
}

func (this *MyQueue) Peek() int {
	this.shift()
	return this.s2.peek()
}

func (this *MyQueue) Empty() bool {
	return this.s1.len()+this.s2.len() == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
//leetcode submit region end(Prohibit modification and deletion)
