package reverselinkedlist

// started at 9:50
// ended at 10:14

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var pP *ListNode
	for head != nil {
		nP := head.Next
		head.Next = pP
		pP = head
		head = nP
	}

	return pP
}
