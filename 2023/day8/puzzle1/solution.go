package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

var step map[rune]int = map[rune]int{
	'L': 0,
	'R': 1,
}

var states = map[string][2]string{}

func main() {
	// file, _ := os.Open("../input.txt")
	file, _ := os.Open("../input-submit.txt")

	steps, _ := Readline(file)
	steps = strings.TrimSpace(steps)
	Readline(file) // Ignore extra spaceful line

	for {
		if line, err := Readline(file); err == nil {
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
