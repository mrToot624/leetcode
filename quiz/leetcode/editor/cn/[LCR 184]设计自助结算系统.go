package main

// leetcode submit region begin(Prohibit modification and deletion)
type Checkout struct {
	mq []int
	q  []int
}

func Constructor() Checkout {
	return Checkout{[]int{}, []int{}}
}

func (this *Checkout) Get_max() int {
	if len(this.mq) == 0 {
		return -1
	}
	return this.mq[0]
}

func (this *Checkout) Add(value int) {
	this.q = append(this.q, value)

	for len(this.mq) > 0 && this.mq[len(this.mq)-1] < value {
		this.mq = this.mq[:len(this.mq)-1]
	}
	this.mq = append(this.mq, value)
}

func (this *Checkout) Remove() int {
	if len(this.q) == 0 {
		return -1
	}
	v := this.q[0]
	this.q = this.q[1:]
	if this.mq[0] == v {
		this.mq = this.mq[1:]
	}
	return v
}

/**
 * Your Checkout object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get_max();
 * obj.Add(value);
 * param_3 := obj.Remove();
 */
//leetcode submit region end(Prohibit modification and deletion)
