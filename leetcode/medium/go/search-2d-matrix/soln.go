package search2dmatrix

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	tPtr, bPtr, lPtr, rPtr := 0, len(matrix)-1, 0, len(matrix[0])-1
	midV := (tPtr + bPtr) / 2
	for tPtr <= bPtr {
		midV = (tPtr + bPtr) / 2
		if matrix[midV][len(matrix[0])-1] >= target && matrix[midV][0] <= target {
			break
		} else if matrix[midV][len(matrix[0])-1] < target {
			tPtr = midV + 1
		} else if matrix[midV][len(matrix[0])-1] > target {
			bPtr = midV - 1
		} else {
			return true
		}
	}

	fmt.Println("Found vertical at: ", midV)
	midH := (lPtr + rPtr) / 2
	for lPtr <= rPtr {
		midH = (lPtr + rPtr) / 2
		if matrix[midV][midH] < target {
			lPtr = midH + 1
		} else if matrix[midV][midH] > target {
			rPtr = midH - 1
		} else {
			return true
		}
	}
	fmt.Println("Found horizoantal at: ", midV)
	return false
}
