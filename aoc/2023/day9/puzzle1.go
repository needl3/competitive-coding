package day9

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/needl3/adventofcode/2023/utils"
)

func Puzzle1() {
	file, _ := os.Open("day9/input.txt")

	sum := 0
	for {
		if line, err := utils.Readline(file); err == nil {
			inp := parseInputLine(line)
			extrapolatedArr := findExtrapolated(inp)
			s := sumExtrapolated(extrapolatedArr)
			sum += s
		} else {
			break
		}
	}
	fmt.Println("Sum: ", sum)
}

func sumExtrapolated(exp [][]int) int {
	s := 0
	prev := 0
	for i := len(exp) - 1; i >= 0; i-- {
		extrapolated := prev + exp[i][len(exp[i])-1]
		s += extrapolated
	}
	return s
}

func findExtrapolated(inp []int) [][]int {
	var exp [][]int
	exp = append(exp, inp)

	for {
		allZeros := true
		var newExp []int
		for i := 0; i < len(exp[len(exp)-1])-1; i++ {
			diff := exp[len(exp)-1][i+1] - exp[len(exp)-1][i]
			allZeros = allZeros && diff == 0
			newExp = append(newExp, diff)
		}

		exp = append(exp, newExp)
		if allZeros {
			break
		}
	}
	return exp
}

func parseInputLine(l string) []int {
	var inp []int
	for _, j := range strings.Split(l, " ") {
		if n, err := strconv.Atoi(j); err == nil {
			inp = append(inp, n)
		}
	}
	return inp
}
