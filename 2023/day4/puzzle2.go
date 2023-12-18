package day4

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Puzzle2() {
	file, _ := os.Open("day4/input.txt")

	sum := 0
	cardMap := map[int]int{}
	for i := 1; ; i++ {
		if line, err := Readline(file); err == nil {
			nMatches := findMatches(line)
			if _, ok := cardMap[i]; ok {
				cardMap[i]++
			} else {
				cardMap[i] = 1
			}
			for j := i + 1; j <= i+nMatches; j++ {
				if _, ok := cardMap[j]; ok {
					cardMap[j] += cardMap[i]
				} else {
					cardMap[j] = cardMap[i]
				}
			}
			sum += cardMap[i]
		} else {
			break
		}
	}
	fmt.Println("Sum: ", sum)
}
func findMatches(line string) int {
	wonMap := map[int]bool{}
	splitted := strings.Split(line, ": ")
	info := strings.Split(splitted[1], " | ")
	won := strings.TrimSpace(info[0])
	have := strings.TrimSpace(info[1])
	for _, w := range strings.Split(won, " ") {
		if n, err := strconv.Atoi(w); err == nil {
			wonMap[n] = true
		}
	}
	matches := 0
	for _, h := range strings.Split(have, " ") {
		if n, err := strconv.Atoi(h); err == nil {
			if _, ok := wonMap[n]; ok {
				matches++
			}
		}
	}
	return matches
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
