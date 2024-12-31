package removeelement

func removeElement(nums []int, val int) int {
	k := 0
	for i := 0; i < len(nums); i++ {
		nums[k] = nums[i]
		if nums[k] != val {
			k++
		}
	}
	return k
}
