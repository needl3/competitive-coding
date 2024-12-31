package numberofstudentsunabletoeatlunch

type Node struct {
	nodeType int
	next     *Node
}

type Queue struct {
	head   *Node
	tail   *Node
	length int
}

func (n *Queue) accept() {
	if n.head != nil {
		n.head = n.head.next
		n.length--
	}
}

func (n *Queue) requeue() {
	if n.head != nil && n.head.next != nil {
		tmp := n.head
		n.head = n.head.next
		n.tail.next = tmp
		n.tail = tmp
		n.tail.next = nil
	}
}

func (n *Queue) traverse() []int {
	arr := []int{}
	for cptr := n.head; cptr != nil; cptr = cptr.next {
		arr = append(arr, cptr.nodeType)
	}
	return arr
}

func countStudents(students []int, sandwiches []int) int {
	studentQ := createQueue(students)
	foodQ := createQueue(sandwiches)

	foodMoved := true
	for foodMoved {
		foodMoved = false
		sessionStudents := studentQ.length

		for i := 0; i < sessionStudents; i++ {
			if studentQ.head.nodeType == foodQ.head.nodeType {
				studentQ.accept()
				foodQ.accept()
				foodMoved = true
			} else {
				studentQ.requeue()
			}
		}
	}
	return studentQ.length
}

func createQueue(arr []int) Queue {
	var head *Node
	var cptr *Node
	for _, v := range arr {
		if head == nil {
			head = &Node{
				nodeType: v,
			}
			cptr = head
		} else {
			cptr.next = &Node{
				nodeType: v,
			}
			cptr = cptr.next
		}
	}
	return Queue{
		head:   head,
		tail:   cptr,
		length: len(arr),
	}
}
