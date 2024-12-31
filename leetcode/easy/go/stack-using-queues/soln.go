package stackusingqueues

import "errors"

type Node struct {
	nodeType int
	next     *Node
}

type Queue struct {
	head *Node
	tail *Node
}

func (n *Queue) dequeue() (int, error) {
	if n.head == nil {
		return 0, errors.New("Cannot dequeue. Queue is empty")
	}

	top := n.head
	n.head = n.head.next
	if n.head == nil {
		n.tail = nil
	}
	return top.nodeType, nil
}

func (n *Queue) queue(val int) {
	newNode := &Node{
		nodeType: val,
	}
	if n.head == nil {
		n.head = newNode
		n.tail = n.head
	} else {
		n.tail.next = newNode
		n.tail = n.tail.next
	}
}

func (n *Queue) back() (int, error) {
	if n.tail == nil {
		return 0, errors.New("Empty queue")
	}
	return n.tail.nodeType, nil
}

type MyStack struct {
	q           Queue
	queueLength int
}

func (s *MyStack) Push(val int) {
	s.q.queue(val)
	s.queueLength++
}

func (s *MyStack) Pop() int {
	for i := 0; i < s.queueLength-1; i++ {
		v, err := s.q.dequeue()
		if err != nil {
			panic(err.Error())
		}
		s.q.queue(v)
	}
	v, err := s.q.dequeue()
	if err != nil {
		panic(err.Error())
	}
	s.queueLength--
	return v
}

func (s *MyStack) Top() int {
	v, err := s.q.back()
	if err != nil {
		panic(err.Error())
	}
	return v
}

func (s *MyStack) Empty() bool {
	_, err := s.q.back()
	if err != nil {
		return true
	}
	return false
}

func Constructor() MyStack {
	return MyStack{}
}
