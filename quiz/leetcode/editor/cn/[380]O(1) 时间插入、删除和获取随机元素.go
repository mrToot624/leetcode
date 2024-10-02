package main

import "math/rand"

// leetcode submit region begin(Prohibit modification and deletion)
type RandomizedSet struct {
	list     []int
	indexMap map[int]int
}

func Constructor_380() RandomizedSet {
	return RandomizedSet{indexMap: make(map[int]int)}
}

func (this *RandomizedSet) Insert(val int) bool {
	_, ok := this.indexMap[val]
	if ok {
		return false
	}
	this.list = append(this.list, val)
	this.indexMap[val] = len(this.list) - 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	curIndex, ok := this.indexMap[val]
	if !ok {
		return false
	}
	last := this.list[len(this.list)-1]
	this.indexMap[last] = curIndex
	this.list[curIndex], this.list[len(this.list)-1] = last, this.list[curIndex]
	this.list = this.list[:len(this.list)-1]
	delete(this.indexMap, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.list[rand.Intn(len(this.list))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
//leetcode submit region end(Prohibit modification and deletion)
