package binarysearch

func search(nums []int, target int) int {
	lptr, rptr := 0, len(nums)-1
	for lptr <= rptr {
		mid := (lptr + rptr) / 2
		if nums[mid] < target {
			lptr = mid + 1
		} else if nums[mid] > target {
			rptr = mid - 1
		} else if nums[mid] == target {
			return mid
		}
	}
	return -1
}
