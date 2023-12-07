package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// file, _ := os.Open("../input.txt")
	file, _ := os.Open("../input-submit.txt")

	var arr []string
	for {
		if line, err := Readline(file); err == nil {
			arr = append(arr, line)
		} else {
			break
		}
	}

	digitLock := false
	digit := 0
	isPartNumberSource := map[string]bool{}
	gearmap := map[string][]int{}
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			if num, err := strconv.Atoi(string(arr[i][j])); err == nil {
				if !digitLock {
					fmt.Printf("Commence %d\nDigit: ", num)
					digitLock = true
				}
				digit = digit*10 + num
				fmt.Printf("%d", num)

				for _, source := range isGear(arr, i, j) {
					isPartNumberSource[source] = true
				}
			} else {
				if digitLock {
					fmt.Printf("\tEnd %d\tIsPart: %v\n", digit, isPartNumberSource)
					for source := range isPartNumberSource {
						gearmap[source] = append(gearmap[source], digit)
					}
					digitLock = false
					digit = 0
					isPartNumberSource = map[string]bool{}
				}
			}
		}
		if digitLock {
			fmt.Printf("\tEnd %d\tIsPart: %v\n", digit, isPartNumberSource)
			for source := range isPartNumberSource {
				gearmap[source] = append(gearmap[source], digit)
			}
			digitLock = false
			digit = 0
			isPartNumberSource = map[string]bool{}
		}
	}

	gearSum := 0
	for _, gear := range gearmap {
		if len(gear) == 2 {
			gearSum += gear[0] * gear[1]
		}
	}
	fmt.Println("Final gear: ", gearSum)
}

func isGear(arr []string, row int, col int) []string {
	var sources []string
	if col > 0 { // left able
		if arr[row][col-1] == '*' { // left
			sources = append(sources, fmt.Sprintf("%v,%v", row, col-1))
		}
		if row > 0 { // top able
			if arr[row-1][col] == '*' { // top
				sources = append(sources, fmt.Sprintf("%v,%v", row-1, col))
			}
			if arr[row-1][col-1] == '*' { // top left
				sources = append(sources, fmt.Sprintf("%v,%v", row-1, col-1))
			}
		}
		if row < len(arr)-1 { // bottom able
			if arr[row+1][col] == '*' { // bottom
				sources = append(sources, fmt.Sprintf("%v,%v", row+1, col))
			}
			if arr[row+1][col-1] == '*' { // bottom left
				sources = append(sources, fmt.Sprintf("%v,%v", row+1, col-1))
			}
		}
	}
	if col < len(arr[row])-1 { // right able
		if arr[row][col+1] == '*' { // right
			sources = append(sources, fmt.Sprintf("%v,%v", row, col+1))
		}
		if row > 0 { // top able
			if arr[row-1][col+1] == '*' { // top right
				sources = append(sources, fmt.Sprintf("%v,%v", row-1, col+1))
			}
		}
		if row < len(arr)-1 { // bottom able
			if arr[row+1][col+1] == '*' { // bottom right
				sources = append(sources, fmt.Sprintf("%v,%v", row+1, col+1))
			}
		}
	}
	return sources
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
