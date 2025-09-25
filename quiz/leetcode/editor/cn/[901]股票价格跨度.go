package main

import "math"

// leetcode submit region begin(Prohibit modification and deletion)
type StockSpanner struct {
	today int
	ms    [][2]int
}

func Constructor_901() StockSpanner {
	return StockSpanner{ms: [][2]int{{math.MaxInt, -1}}}
}

func (this *StockSpanner) Next(price int) int {
	for len(this.ms) > 0 && this.ms[len(this.ms)-1][0] <= price {
		this.ms = this.ms[:len(this.ms)-1]
	}
	span := this.today - this.ms[len(this.ms)-1][1]
	this.ms = append(this.ms, [2]int{price, this.today})
	this.today++
	return span
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */
//leetcode submit region end(Prohibit modification and deletion)
