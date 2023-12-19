package day14

import (
	"fmt"
	"os"

	"github.com/needl3/adventofcode/2023/utils"
)

func Puzzle1() {
	file, _ := os.Open("day14/input.txt")

	parsedInput := parseInput(file)

	fmt.Println("Sum: ", findSum(parsedInput))
}

func findSum(parsedInput []string) int {
	sum := 0
	for y := 0; y < len(parsedInput[0]); y++ {
		lastPos := len(parsedInput)
		for x := 0; x < len(parsedInput); x++ {
			if parsedInput[x][y] == '#' {
				lastPos = len(parsedInput) - x - 1
			} else if parsedInput[x][y] == 'O' {
				sum += lastPos
				lastPos--
			}
		}
	}
	return sum
}

func parseInput(file *os.File) []string {
	parsed := []string{}
	for {
		if line, err := utils.Readline(file); err == nil {
			parsed = append(parsed, line)
		} else {
			break
		}
	}
	return parsed
}
