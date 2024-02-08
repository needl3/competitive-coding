package day2

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/needl3/adventofcode/2023/utils"
)

var Bag = map[string]int{
	"r": 12,
	"g": 13,
	"b": 14,
}

func Puzzle1() {
	file, _ := os.Open("day2/input.txt")
	defer file.Close()

	sum := 0
	for {
		if line, err := utils.Readline(file); err == nil {
			if valid, id := IsValid(line); valid {
				sum += id
			}
		} else {
			break
		}
	}
	fmt.Printf("Final ans: %v\n", sum)
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
