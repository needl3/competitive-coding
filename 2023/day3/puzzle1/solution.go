package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("../input-submit.txt")

	var arr []string
	for {
		if line, err := Readline(file); err == nil {
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
					fmt.Printf("Commence %d\nDigit: ", num)
					digitLock = true
				}
				digit = digit*10 + num
				fmt.Printf("%d", num)
				isPartNumber = isPartNumber || isSymbol(arr, i, j)
			} else {
				if digitLock {
					fmt.Printf("\tEnd %d\tIsPart: %v\n", digit, isPartNumber)
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
			fmt.Printf("\tEnd %d\tIsPart: %v\n", digit, isPartNumber)
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
		if !isDigit(arr[row][col-1]) && arr[row][col-1] != '.' { // left
			return true
		}
		if row > 0 { // top able
			if !isDigit(arr[row-1][col]) && arr[row-1][col] != '.' { // top
				return true
			}
			if !isDigit(arr[row-1][col-1]) && arr[row-1][col-1] != '.' { // top left
				return true
			}
		}
		if row < len(arr)-1 { // bottom able
			if !isDigit(arr[row+1][col]) && arr[row+1][col] != '.' { // bottom
				return true
			}
			if !isDigit(arr[row+1][col-1]) && arr[row+1][col-1] != '.' { // bottom left
				return true
			}
		}
	}
	if col < len(arr[row])-1 { // right able
		if !isDigit(arr[row][col+1]) && arr[row][col+1] != '.' { // right
			return true
		}
		if row > 0 { // top able
			if !isDigit(arr[row-1][col+1]) && arr[row-1][col+1] != '.' { // top right
				return true
			}
		}
		if row < len(arr)-1 { // bottom able
			if !isDigit(arr[row+1][col+1]) && arr[row+1][col+1] != '.' { // bottom right
				return true
			}
		}
	}
	return false
}

func isDigit(a byte) bool {
	s := rune(a)
	return s == '1' || s == '2' || s == '3' || s == '4' || s == '5' || s == '6' || s == '7' || s == '8' || s == '9' || s == '0'
}

func Readline(f *os.File) (string, error) {
	var line []byte
	singleFileContent := make([]byte, 1)
	i := 0
	for {
		if read, err := f.Read(singleFileContent); err == nil && read > 0 {
			if singleFileContent[0] != byte('\n') {
				line = append(line, singleFileContent...)
			} else {
				break
			}
		} else {
			return "", errors.New("eof")
		}
		i = i + 1
	}
	return string(line), nil
}

