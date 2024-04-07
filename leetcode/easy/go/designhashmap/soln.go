package designhashmap

import "fmt"

var HASH_SIZE = 100

type MyHashNode struct {
	key  int
	val  int
	next *MyHashNode
}

type MyHashMap struct {
	set []*MyHashNode
}

func Constructor() MyHashMap {
	return MyHashMap{
		set: make([]*MyHashNode, HASH_SIZE),
	}
}

func (this *MyHashMap) Hash(key int) int{
    return 7919 * key % HASH_SIZE
}

func (this *MyHashMap) Put(key int, value int) {
	if ptr := this.set[this.Hash(key)]; ptr != nil {
		for ; ptr != nil; ptr = ptr.next {
			if ptr.key == key {
				ptr.val = value
				break
			}
			if ptr.next == nil {
				ptr.next = &MyHashNode{
					key: key,
					val: value,
				}
			}
		}
	} else {
		this.set[this.Hash(key)] = &MyHashNode{
			key: key,
			val: value,
		}
	}
}

func (this *MyHashMap) Remove(key int) {
	ptr := this.set[this.Hash(key)]
	if ptr != nil && ptr.key == key {
		this.set[this.Hash(key)] = ptr.next
	} else {
		for ; ptr != nil; ptr = ptr.next {
			if ptr.next != nil && ptr.next.key == key {
				ptr.next = ptr.next.next
				break
			}
		}
	}
}

func (this *MyHashMap) Get(key int) int {
	ptr := this.set[this.Hash(key)]

	for ; ptr != nil; ptr = ptr.next {
		if ptr.key == key {
			return ptr.val
		}
	}
	return -1
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */

func Exec() {
	obj := Constructor()
	obj.Put(1, 1)
	obj.Put(2, 2)
	fmt.Println(obj.Get(1))
	fmt.Println(obj.Get(3))
	obj.Put(2, 1)
	fmt.Println(obj.Get(2))
	obj.Remove(2)
	fmt.Println(obj.Get(2))
	fmt.Println(obj)
}
