package removeduplicatesfromsortedarray

func RemoveDuplicates(nums []int) int {
	if len(nums) == 1 {
		return 1
	}

	i := 0
	j := 1
	uniqueCount := 1
	for j < len(nums) {
		if nums[j] == nums[i] {
			j++
		} else {
			i++
			nums[i] = nums[j]
			j++
			uniqueCount++
		}
	}

	return uniqueCount
}
