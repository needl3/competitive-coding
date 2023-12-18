package day8

import (
	"fmt"
	"os"
	"strings"

	"github.com/needl3/adventofcode/2023/utils"
)

var step map[rune]int = map[rune]int{
	'L': 0,
	'R': 1,
}

var states = map[string][2]string{}

func Puzzle1() {
	file, _ := os.Open("day8/input.txt")

	steps, _ := utils.Readline(file)
	steps = strings.TrimSpace(steps)
	utils.Readline(file) // Ignore extra spaceful line

	for {
		if line, err := utils.Readline(file); err == nil {
			data := strings.Split(line, " = ")
			state := data[0]
			children := strings.Split(data[1], ", ")
			lChild := children[0][1:]
			rChild := children[1][:len(children[1])-1]

			states[state] = [2]string{lChild, rChild}
		} else {
			break
		}
	}

	count := 0
	head := "AAA"
	for ; ; count++ {
		if head == "ZZZ" {
			break
		} else {
			ind := step[rune(steps[count%len(steps)])]
			head = states[head][ind]
		}
	}
	fmt.Println(count)
}
