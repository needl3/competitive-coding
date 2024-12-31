package mergetwosortedlists

import (
	"fmt"
	"slices"
	"testing"
)

func TestSoln1(t *testing.T) {
	l1 := ListNode{
		Val: 11,
		Next: &ListNode{
			Val: 20,
			Next: &ListNode{
				Val: 30,
			},
		},
	}

	l2 := ListNode{
		Val: 3,
		Next: &ListNode{
			Val: 12,
			Next: &ListNode{
				Val: 25,
			},
		},
	}

	merged := mergeTwoLists(&l1, &l2)
	var vals []int
	for merged != nil {
		vals = append(vals, merged.Val)
		merged = merged.Next
	}

	fmt.Println("Merged: ", vals)
	if !slices.Equal([]int{3, 11, 12, 20, 25, 30}, vals) {
		t.Fail()
	}
}
