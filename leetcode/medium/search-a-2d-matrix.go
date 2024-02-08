func searchMatrix(matrix [][]int, target int) bool {
	if rowIndex := binSearchRow(matrix, target); rowIndex >= 0 {
		foundIndex := binSearch(matrix[rowIndex], target)
		return foundIndex >= 0
	}
	return false
}

func binSearchRow(matrix [][]int, target int) int {
	lptr := 0
	rptr := len(matrix) - 1
	for lptr <= rptr {
		mid := (lptr + rptr) / 2
		if target > matrix[mid][len(matrix[mid])-1] {
			lptr = mid + 1
		} else if target < matrix[mid][0] {
			rptr = mid - 1
		} else {
			return mid
		}
	}
	return -1
}

func binSearch(row []int, target int) int {
	lptr := 0
	rptr := len(row)
	for lptr <= rptr {
		mid := (lptr + rptr) / 2
		if target > row[mid] {
			lptr = mid + 1
		} else if target < row[mid] {
			rptr = mid - 1
		} else {
			return mid
		}
	}
	return -1
}
