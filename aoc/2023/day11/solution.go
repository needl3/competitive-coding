package day11

import (
	"os"

	"github.com/needl3/adventofcode/2023/utils"
)

func Puzzle1() {
	file, _ := os.Open("day11/test.txt")

	var input []string
	for {
		if line, err := utils.Readline(file); err == nil {
			input = append(input, line)
		} else {
			break
		}
	}

}
