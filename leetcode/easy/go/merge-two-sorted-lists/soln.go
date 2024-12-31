package mergetwosortedlists

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var head, top *ListNode
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			assignHead(&head, &list1)
		} else {
			assignHead(&head, &list2)
		}
		if top == nil {
			top = head
		}
	}

	for list1 != nil {
		assignHead(&head, &list1)
	}
	for list2 != nil {
		assignHead(&head, &list2)
	}

	return top
}

func assignHead(head **ListNode, minNode **ListNode) {
	if (*head) == nil {
		*head = *minNode
	} else {
		(*head).Next = *minNode
		*head = *minNode
	}
	*minNode = (*minNode).Next
}
