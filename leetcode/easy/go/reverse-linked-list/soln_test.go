package reverselinkedlist

import "testing"

func TestSoln1(t *testing.T) {
	head := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 3,
		},
	}

	head = reverseList(head)
	if head.Val != 3 && head.Next.Val != 2 {
		t.Fail()
	}
}
