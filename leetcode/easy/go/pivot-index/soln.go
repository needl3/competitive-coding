package pivotindex

import "fmt"

func pivotIndex(nums []int) int {
	cumFront := make([]int, len(nums))
	for i, frontVal := range nums {
		if i == 0 {
			cumFront[i] = frontVal
			continue
		}

		cumFront[i] = cumFront[i-1] + frontVal
	}
	cumBack := 0
	for i := len(nums) - 1; i >= 0; i-- {
		cumBack += nums[i]
		if cumFront[len(nums)-1-i] == cumBack {
			return len(nums) - 1 - i
		}
	}
	return -1
}

func Exec() {
	fmt.Println(pivotIndex([]int{1, 7, 3, 6, 5, 6}))
}
