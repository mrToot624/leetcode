package main

// leetcode submit region begin(Prohibit modification and deletion)
type LFUListNode struct {
	k, v, freq int
	next, prev *LFUListNode
}

type LFULinkedList struct {
	head, tail *LFUListNode
	l          int
}

func (list *LFULinkedList) addToTail(node *LFUListNode) {
	node.prev = list.tail.prev
	node.next = list.tail
	list.tail.prev.next = node
	list.tail.prev = node
	list.l++
}

func (list *LFULinkedList) remove(node *LFUListNode) int {
	node.next.prev = node.prev
	node.prev.next = node.next
	list.l--
	return node.k
}

type LFUCache struct {
	cap, minFreq   int
	cache          map[int]*LFUListNode
	freqLinkedList map[int]*LFULinkedList
}

func Constructor_LFU(capacity int) LFUCache {
	return LFUCache{
		cap:            capacity,
		cache:          make(map[int]*LFUListNode),
		freqLinkedList: make(map[int]*LFULinkedList),
	}
}

func (this *LFUCache) increaseFreq(node *LFUListNode) {
	curFreq := node.freq
	this.freqLinkedList[curFreq].remove(node)
	this.cleanupFreqLinkedList(curFreq)

	node.freq++
	this.addNewFreqListIfNotPresent(node.freq).addToTail(node)
}

func (this *LFUCache) removeLeastFreq() int {
	minFreqLinkedList := this.freqLinkedList[this.minFreq]
	key := minFreqLinkedList.remove(minFreqLinkedList.head.next)
	this.cleanupFreqLinkedList(this.minFreq)
	return key
}

func (this *LFUCache) insert(key, value int) *LFUListNode {
	node := &LFUListNode{k: key, v: value, freq: 1}
	this.addNewFreqListIfNotPresent(1).addToTail(node)
	return node
}

func (this *LFUCache) addNewFreqListIfNotPresent(freq int) *LFULinkedList {
	_, ok := this.freqLinkedList[freq]
	if !ok {
		head, tail := &LFUListNode{}, &LFUListNode{}
		head.next = tail
		tail.prev = head
		this.freqLinkedList[freq] = &LFULinkedList{head: head, tail: tail}
	}
	if freq == 1 || freq < this.minFreq {
		this.minFreq = freq
	}
	return this.freqLinkedList[freq]
}

func (this *LFUCache) cleanupFreqLinkedList(freq int) {
	if this.freqLinkedList[freq].l == 0 {
		delete(this.freqLinkedList, freq)
		if this.minFreq == freq {
			this.minFreq++
		}
	}
}

func (this *LFUCache) Get(key int) int {
	node, ok := this.cache[key]
	if !ok {
		return -1
	}

	this.increaseFreq(node)
	return node.v
}

func (this *LFUCache) Put(key int, value int) {
	node, ok := this.cache[key]
	if ok {
		node.v = value
		this.increaseFreq(node)
		return
	}

	if len(this.cache) == this.cap {
		delete(this.cache, this.removeLeastFreq())
	}
	this.cache[key] = this.insert(key, value)
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
//leetcode submit region end(Prohibit modification and deletion)
