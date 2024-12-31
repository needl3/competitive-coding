package removeelement

import (
	"log"
	"testing"
)

func TestSoln1(t *testing.T) {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2
	count := removeElement(nums, val)

	log.Println("Count: ", count)
	log.Println("Nums: ", nums)

	if count != 5 {
		t.Fail()
	}

	answerArray := []int{0, 1, 3, 0, 4}
	for i, ans := range answerArray {
		if nums[i] != ans {
			t.Fail()
		}
	}

}

func TestSoln2(t *testing.T) {
	nums := []int{1}
	val := 1
	count := removeElement(nums, val)

	log.Println("Count: ", count)
	log.Println("Nums: ", nums)

	if count != 0 {
		t.Fail()
	}
}

func TestSoln3(t *testing.T) {
	nums := []int{2, 2, 3}
	val := 2
	count := removeElement(nums, val)

	log.Println("Count: ", count)
	log.Println("Nums: ", nums)

	if count != 1 {
		t.Fail()
	}
}
