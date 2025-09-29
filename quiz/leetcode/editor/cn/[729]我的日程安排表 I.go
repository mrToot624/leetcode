package main

import (
	"github.com/emirpasic/gods/trees/redblacktree"
)

// leetcode submit region begin(Prohibit modification and deletion)
type MyCalendar struct {
	rbt *redblacktree.Tree
}

func Constructor() MyCalendar {
	return MyCalendar{redblacktree.NewWithIntComparator()}
}

func (this *MyCalendar) Book(startTime int, endTime int) bool {
	if ceiling, ok := this.rbt.Ceiling(startTime); ok {
		if ceiling.Key.(int) < endTime {
			return false
		}
	}
	if floor, ok := this.rbt.Floor(startTime); ok {
		if floor.Value.(int) > startTime {
			return false
		}
	}
	this.rbt.Put(startTime, endTime)
	return true
}

/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(startTime,endTime);
 */
//leetcode submit region end(Prohibit modification and deletion)
