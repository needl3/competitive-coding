package day9

import (
	"fmt"
	"os"

	"github.com/needl3/adventofcode/2023/utils"
)

func Puzzle2() {
	file, _ := os.Open("day9/input.txt")

	sum := 0
	for {
		if line, err := utils.Readline(file); err == nil {
			inp := parseInputLine(line)
			extrapolatedArr := findExtrapolated(inp)
			s := sumExtrapolatedFront(extrapolatedArr)
			sum += s
		} else {
			break
		}
	}
	fmt.Println("Sum: ", sum)
}

func sumExtrapolatedFront(exp [][]int) int {
	extrapolated := 0
	for i := len(exp) - 1; i >= 0; i-- {
		extrapolated = exp[i][0] - extrapolated
	}
	return extrapolated
}
