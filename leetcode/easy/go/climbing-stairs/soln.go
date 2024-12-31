package climbingstairs

func climbStairs(n int) int {
	return memoizedRecursion(n, &map[int]int{})
}

func memoizedRecursion(n int, memo *map[int]int) int {
	if n == 0 {
		return 1
	}
	if n < 0 {
		return 0
	}
	n2, ok := (*memo)[n-2]
	if !ok {
		n2 = memoizedRecursion(n-2, memo)
		(*memo)[n-2] = n2
	}
	n1, ok := (*memo)[n-1]
	if !ok {
		n1 = memoizedRecursion(n-1, memo)
		(*memo)[n-1] = n1
	}
	return n1 + n2
}
