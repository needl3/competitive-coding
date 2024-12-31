package binarysearch

import "testing"

func TestSoln(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	v := search(arr, 4)
	t.Log("Position: ", v)
	if v != 4 {
		t.Fail()
	}
}
