package day14

import (
	"fmt"
	"os"

	"github.com/needl3/adventofcode/2023/utils"
)

func printMat(inp [][]rune) {
	for _, i := range inp {
		for _, j := range i {
			fmt.Printf("%s", string(j))
		}
		fmt.Println()
	}
	fmt.Println()
}

const CYCLES = 1000000000

func Puzzle2() {
	file, _ := os.Open("day14/input.txt")
	parsedInput := parseInputRune(file)

	repeatMap := map[string]int{}
	for i := 0; i < CYCLES; i++ {
		tiltNorth(parsedInput)
		tiltWest(parsedInput)
		tiltSouth(parsedInput)
		tiltEast(parsedInput)

		stringified := stringify(parsedInput)
		if index, ok := repeatMap[stringified]; ok {
			i = CYCLES - (CYCLES-index)%(i-index)
		} else {
			repeatMap[stringified] = i
		}
	}
	fmt.Println("Sum: ", findSumRune(parsedInput))
}

func stringify(inp [][]rune) string {
	s := ""
	for _, i := range inp {
		for _, j := range i {
			s += string(j)
		}
	}
	return s
}

func findSumRune(parsedInput [][]rune) int {
	sum := 0
	for y := 0; y < len(parsedInput[0]); y++ {
		for x := 0; x < len(parsedInput); x++ {
			if parsedInput[x][y] == 'O' {
				sum += len(parsedInput) - x
			}
		}
	}
	return sum
}

func parseInputRune(file *os.File) [][]rune {
	out := [][]rune{}
	for line, err := utils.Readline(file); err == nil; line, err = utils.Readline(file) {
		out = append(out, make([]rune, len(line)))
		for i, c := range line {
			out[len(out)-1][i] = c
		}
	}
	return out
}

func tiltNorth(parsedInput [][]rune) {
	for y := 0; y < len(parsedInput[0]); y++ {
		lastPos := 0
		for x := 0; x < len(parsedInput); x++ {
			if parsedInput[x][y] == '#' {
				lastPos = x + 1
			} else if parsedInput[x][y] == 'O' {
				parsedInput[x][y] = '.'
				parsedInput[lastPos][y] = 'O'
				lastPos++
			}
		}
	}
}

func tiltSouth(parsedInput [][]rune) {
	for y := 0; y < len(parsedInput[0]); y++ {
		lastPos := len(parsedInput) - 1
		for x := len(parsedInput) - 1; x >= 0; x-- {
			if parsedInput[x][y] == '#' {
				lastPos = x - 1
			} else if parsedInput[x][y] == 'O' {
				parsedInput[x][y] = '.'
				parsedInput[lastPos][y] = 'O'
				lastPos--
			}
		}
	}
}

func tiltWest(parsedInput [][]rune) {
	for x := 0; x < len(parsedInput); x++ {
		lastPos := 0
		for y := 0; y < len(parsedInput[0]); y++ {
			if parsedInput[x][y] == '#' {
				lastPos = y + 1
			} else if parsedInput[x][y] == 'O' {
				parsedInput[x][y] = '.'
				parsedInput[x][lastPos] = 'O'
				lastPos++
			}
		}
	}
}

func tiltEast(parsedInput [][]rune) {
	for x := 0; x < len(parsedInput); x++ {
		lastPos := len(parsedInput[0]) - 1
		for y := len(parsedInput[0]) - 1; y >= 0; y-- {
			if parsedInput[x][y] == '#' {
				lastPos = y - 1
			} else if parsedInput[x][y] == 'O' {
				parsedInput[x][y] = '.'
				parsedInput[x][lastPos] = 'O'
				lastPos--
			}
		}
	}
}
