package day3

import (
	"fmt"
	"os"
	"strconv"

	"github.com/needl3/adventofcode/2023/utils"
)

func Puzzle2() {
	file, _ := os.Open("day3/input.txt")

	var arr []string
	for {
		if line, err := utils.Readline(file); err == nil {
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
					digitLock = true
				}
				digit = digit*10 + num

				for _, source := range isGear(arr, i, j) {
					isPartNumberSource[source] = true
				}
			} else {
				if digitLock {
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
