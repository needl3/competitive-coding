package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var Bag = map[string]int{
	"r": 12,
	"g": 13,
	"b": 14,
}

func main() {
	file, _ := os.Open("../input-submit.txt")
	defer file.Close()

	sum := 0
	for {
		if line, err := Readline(file); err == nil {
			if valid, id := IsValid(line); valid {
				sum += id
			}
		} else {
			break
		}
	}
	fmt.Printf("Final ans: %v\n", sum)
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

func IsValid(line string) (bool, int) {
	splitted := strings.Split(line, ": ")
	valid := true

	playArr := strings.Split(splitted[1], "; ")
	for _, play := range playArr {
		set := strings.Split(play, ", ")
		for _, col := range set {
			item := strings.Split(col, " ")
			if maxN, ok := Bag[string(item[1][0])]; ok {
				if fromBag, err := strconv.Atoi(item[0]); err == nil && fromBag > maxN {
					valid = false
					break
				}
			}
		}
		if !valid {
			break
		}
	}
	id, _ := strconv.Atoi(string(strings.Split(splitted[0], " ")[1]))
	return valid, id
}
