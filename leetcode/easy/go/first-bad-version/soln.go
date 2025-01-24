package firstbadversion

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func isBadVersion(n int) bool {
	if n <= 0 {
		return false
	}
	return true
}

func firstBadVersion(n int) int {
	lptr, rptr := 0, n
	prevMid := -1
	for lptr < rptr {
		mid := (lptr + rptr) / 2
		if isBadVersion(mid) {
			if prevMid == mid {
				return mid
			}
			rptr = mid
		} else {
			if prevMid == mid {
				return mid + 1
			}
			lptr = mid
		}
		prevMid = mid
	}
	return -1
}
