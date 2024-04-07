package designhashset

import "fmt"

var HASH_SIZE = 10000

type MyHashNode struct {
	key  int
	next *MyHashNode
}

type MyHashSet struct {
	set []*MyHashNode
}

func Constructor() MyHashSet {
	return MyHashSet{
		set: make([]*MyHashNode, HASH_SIZE),
	}
}

func (this *MyHashSet) Add(key int) {
  if ptr := this.set[key%HASH_SIZE]; ptr != nil {
    for ; ptr != nil; ptr = ptr.next {
        if ptr.key == key{
          return
        }
        if ptr.next == nil{
          ptr.next = &MyHashNode{
            key: key,
          }
        }
    }
  } else {
    this.set[key % HASH_SIZE] = &MyHashNode{
      key: key,
    }
  }
}

func (this *MyHashSet) Remove(key int) {	
  ptr := this.set[key % HASH_SIZE]
  if ptr != nil && ptr.key == key{
      this.set[key % HASH_SIZE] = ptr.next
  } else {
    for ; ptr != nil; ptr = ptr.next{
      if ptr.next != nil && ptr.next.key == key{
          ptr.next = ptr.next.next
          break
      }
    }
  }
}

func (this *MyHashSet) Contains(key int) bool {
	ptr := this.set[key%HASH_SIZE]

	for ; ptr != nil; ptr = ptr.next {
		if ptr.key == key {
			return true
		}
	}
	return false
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
	obj.Add(1)
	obj.Add(2)
	fmt.Println(obj.Contains(1))
	fmt.Println(obj.Contains(3))
	obj.Add(2)
	fmt.Println(obj.Contains(2))
	obj.Remove(2)
	fmt.Println(obj.Contains(2))
        fmt.Println(obj)
}
