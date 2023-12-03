package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatalf("Unable to open file\n")
	}

	sum := 0
	first := 0
	last := 0
	isFirst := true

	buf := make([]byte, 1)

	var eof error
	for eof == nil {
		_, eof = file.Read(buf)
		if buf[0] == byte('\n') {
			if finalDigit, err := strconv.Atoi(fmt.Sprintf("%v%v", first, last)); err == nil {
				sum = sum + finalDigit
			}
			first = 0
			last = 0
			isFirst = true
		}
		if data, err := strconv.Atoi(string(buf)); err == nil {
			if isFirst {
				first = data
				isFirst = false
			}
			last = data
		}
	}
	fmt.Printf("Final answer: %v\n", sum)
}
