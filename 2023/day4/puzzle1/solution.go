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
		wonMap := map[int]bool{}
		if line, err := Readline(file); err == nil {
			info := strings.Split(strings.Split(line, ": ")[1], " | ")
			won := strings.TrimSpace(info[0])
			have := strings.TrimSpace(info[1])
			for _, w := range strings.Split(won, " ") {
				if n, err := strconv.Atoi(w); err == nil {
					wonMap[n] = true
				}
			}
			point := 0
			for _, h := range strings.Split(have, " ") {
				if n, err := strconv.Atoi(h); err == nil {
					if _, ok := wonMap[n]; ok {
						if point == 0 {
							point = 1
						} else {
							point *= 2
						}
					}
				}
			}
			sum += point
		} else {
			break
		}
	}
	fmt.Println("Final sum: ", sum)
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
