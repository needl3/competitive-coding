package search2dmatrix

import "testing"

func TestSoln(t *testing.T) {
	arr := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}

	present := searchMatrix(arr, 10)
	if !present {
		t.Fail()
	}
}
