package groupanagrams

import "fmt"

type LRUCacheNode struct {
	value int
	key   int
	next  *LRUCacheNode
	prev  *LRUCacheNode
}
type LRUCache struct {
	cache    map[int]*LRUCacheNode
	capacity int
	front    *LRUCacheNode
	back     *LRUCacheNode
	length   int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		cache:    make(map[int]*LRUCacheNode, capacity),
		capacity: capacity,
		length:   0,
	}
}

func (this *LRUCache) Get(key int) int {
	fmt.Println("Getting ", key)
	if node, ok := this.cache[key]; ok && node != nil {
		this.MarkRecent(node)
		return node.value
	}

	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.cache[key]; ok {
		node.value = value
		this.MarkRecent(node)
		return
	}

	fmt.Println("Putting ", key, value, " length=", this.length, len(this.cache))
	if this.length >= this.capacity {
		this.RemoveLRU()
	}
	this.AddToHead(key, value)
}

func (this *LRUCache) Traverse() {
	for node := this.front; node != nil; node = node.prev {
		fmt.Print(node.key, " => ")
	}
	fmt.Println()
}

func (this *LRUCache) AddToHead(key int, val int) {
	tempNode := this.front
	this.front = &LRUCacheNode{
		value: val,
		key:   key,
		prev:  tempNode,
	}
	if tempNode == nil {
		this.back = this.front
	} else {
		tempNode.next = this.front
	}
	this.cache[key] = this.front
	this.length++
}

func (this *LRUCache) RemoveLRU() {
	fmt.Println("Removing LRU for length ", this.length)
	if this.length == 0 {
		return
	} else if this.length == 1 {
		delete(this.cache, this.back.key)
		this.front = nil
		this.back = nil
	} else {
		delete(this.cache, this.back.key)
		this.back = this.back.next
		this.back.prev = nil
	}
	this.length--
	// Garbage collector will free this.back
}

func (this *LRUCache) MarkRecent(node *LRUCacheNode) {
	fmt.Println("Marking recent")
	if node == this.front {
		return
	}

	if node == this.back {
		this.back = node.next
		this.back.prev = nil
	} else {
		node.next.prev = node.prev
		node.prev.next = node.next
	}

	this.front.next = node
	node.prev = this.front
	node.next = nil

	this.front = node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func Exec() {
	obj := Constructor(2)
	obj.Put(1, 1)
	obj.Put(2, 2)
	obj.Get(1)
	obj.Put(3, 3)
	obj.Get(2)
	obj.Traverse()
	obj.Put(4, 4)
	obj.Get(1)

	obj.Traverse()
}
