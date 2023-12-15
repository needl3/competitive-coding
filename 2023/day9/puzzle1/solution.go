package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// file, _ := os.Open("../input.txt")
	file, _ := os.Open("../input-submit.txt")

	sum := 0
	for {
		if line, err := Readline(file); err == nil {
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
