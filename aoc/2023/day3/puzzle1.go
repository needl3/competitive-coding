package day3

import (
	"fmt"
	"os"
	"strconv"

	"github.com/needl3/adventofcode/2023/utils"
)

func Puzzle1() {
	file, _ := os.Open("day3/input.txt")

	var arr []string
	for {
		if line, err := utils.Readline(file); err == nil {
			arr = append(arr, line)
		} else {
			break
		}
	}

	sum := 0
	digitLock := false
	digit := 0
	isPartNumber := false
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			if num, err := strconv.Atoi(string(arr[i][j])); err == nil {
				if !digitLock {
					digitLock = true
				}
				digit = digit*10 + num
				isPartNumber = isPartNumber || isSymbol(arr, i, j)
			} else {
				if digitLock {
					if isPartNumber {
						sum += digit
					}
					digitLock = false
					digit = 0
					isPartNumber = false
				}
			}
		}
		if digitLock {
			if isPartNumber {
				sum += digit
			}
			digitLock = false
			digit = 0
			isPartNumber = false
		}
	}
	fmt.Printf("Final answer: %d", sum)
}

func isSymbol(arr []string, row int, col int) bool {
	if col > 0 { // left able
		if !utils.IsDigit(arr[row][col-1]) && arr[row][col-1] != '.' { // left
			return true
		}
		if row > 0 { // top able
			if !utils.IsDigit(arr[row-1][col]) && arr[row-1][col] != '.' { // top
				return true
			}
			if !utils.IsDigit(arr[row-1][col-1]) && arr[row-1][col-1] != '.' { // top left
				return true
			}
		}
		if row < len(arr)-1 { // bottom able
			if !utils.IsDigit(arr[row+1][col]) && arr[row+1][col] != '.' { // bottom
				return true
			}
			if !utils.IsDigit(arr[row+1][col-1]) && arr[row+1][col-1] != '.' { // bottom left
				return true
			}
		}
	}
	if col < len(arr[row])-1 { // right able
		if !utils.IsDigit(arr[row][col+1]) && arr[row][col+1] != '.' { // right
			return true
		}
		if row > 0 { // top able
			if !utils.IsDigit(arr[row-1][col+1]) && arr[row-1][col+1] != '.' { // top right
				return true
			}
		}
		if row < len(arr)-1 { // bottom able
			if !utils.IsDigit(arr[row+1][col+1]) && arr[row+1][col+1] != '.' { // bottom right
				return true
			}
		}
	}
	return false
}
