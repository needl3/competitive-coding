package main

import (
	"errors"
	"log"
	"os"
)

var strToNumMap = map[string]int{
	"zero":  0,
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"0":     0,
	"1":     1,
	"2":     2,
	"3":     3,
	"4":     4,
	"5":     5,
	"6":     6,
	"7":     7,
	"8":     8,
	"9":     9,
}

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatalf("Unable to open file\n")
	}
	defer file.Close()

	sum := 0
	for {
		if line, err := readline(file); err == nil {
			num := findNum(line)
			log.Printf("%v: %v\n", line, num)
			sum = sum + num
		} else {
			break
		}
	}
	log.Printf("Final answer: %v\n", sum)
}

func findNum(line string) int {
	first := -1
	firstPos := -1
	last := -1
	lastPos := -1

	for key, val := range strToNumMap {
		begin := 0
		for begin+len(key) <= len(line) {
			if line[begin:begin+len(key)] == key {
				if firstPos > begin+len(key) || first == -1 {
					first = val
					firstPos = begin + len(key)
				}
				if lastPos < begin+len(key) {
					last = val
					lastPos = begin + len(key)
				}
			}
			begin = begin + 1
		}
	}
	return first*10 + last
}

func readline(f *os.File) (string, error) { // offset, string
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
