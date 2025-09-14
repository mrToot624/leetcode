package main

// leetcode submit region begin(Prohibit modification and deletion)
type LRUListNode struct {
	k, v       int
	next, prev *LRUListNode
}

type LRULinkedList struct {
	head, tail *LRUListNode
}

func (l *LRULinkedList) refresh(node *LRUListNode) {
	node.next.prev = node.prev
	node.prev.next = node.next
	l.addToTail(node)
}

func (l *LRULinkedList) insert(key, value int) *LRUListNode {
	node := &LRUListNode{k: key, v: value}
	l.addToTail(node)
	return node
}

func (l *LRULinkedList) addToTail(node *LRUListNode) {
	node.next = l.tail
	node.prev = l.tail.prev
	l.tail.prev.next = node
	l.tail.prev = node
}

func (l *LRULinkedList) removeOldest() int {
	key := l.head.next.k
	l.head.next = l.head.next.next
	l.head.next.prev = l.head
	return key
}

type LRUCache struct {
	cap   int
	cache map[int]*LRUListNode
	list  LRULinkedList
}

func Constructor_LRU(capacity int) LRUCache {
	head, tail := &LRUListNode{}, &LRUListNode{}
	head.next = tail
	tail.prev = head
	return LRUCache{
		cap:   capacity,
		cache: make(map[int]*LRUListNode),
		list:  LRULinkedList{head, tail},
	}
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.cache[key]
	if !ok {
		return -1
	}

	this.list.refresh(node)
	return node.v
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.cache[key]
	if ok {
		node.v = value
		this.list.refresh(node)
		return
	}

	if len(this.cache) == this.cap {
		delete(this.cache, this.list.removeOldest())
	}
	this.cache[key] = this.list.insert(key, value)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
//leetcode submit region end(Prohibit modification and deletion)
