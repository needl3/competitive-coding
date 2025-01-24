package guessnumberhigherorlower

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessFactory(num int) func(n int) int {
	return func(n int) int {
		if n > num {
			return -1
		} else if n < num {
			return 1
		}
		return 0
	}
}

func guessNumber(n int) int {
	guess := guessFactory(889)
	lptr, rptr := 0, n
	for {
		mid := (lptr + rptr) / 2
		m := guess(mid)
		if m < 0 {
			rptr = mid - 1
		} else if m > 0 {
			lptr = mid + 1
		} else if m == 0 {
			return mid
		}
	}
}
