package removeduplicatesfromsortedarray

import (
	"log"
	"testing"
)

func TestSoln(t *testing.T) {
	answer := []int{1, 2, 2}
	nums := []int{1, 1, 2}
	uniqueValues := RemoveDuplicates(nums)
	if uniqueValues != 2 {
		t.Fail()
	}
	log.Println(nums)
	for i, x := range nums {
		if answer[i] != x {
			t.Fail()
		}
	}
}
