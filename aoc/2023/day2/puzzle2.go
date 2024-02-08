package day2

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/needl3/adventofcode/2023/utils"
)

func Puzzle2() {
	file, _ := os.Open("day2/input.txt")
	defer file.Close()

	sum := 0
	for {
		if line, err := utils.Readline(file); err == nil {
			if valid, cube := IsValidV2(line); valid {
				sum += cube
			}
		} else {
			break
		}
	}
	fmt.Printf("Final ans: %v\n", sum)
}

func IsValidV2(line string) (bool, int) {
	splitted := strings.Split(line, ": ")
	valid := true
	mapCube := map[string]int{
		"r": 0,
		"b": 0,
		"g": 0,
	}

	playArr := strings.Split(splitted[1], "; ")
	for _, play := range playArr {
		set := strings.Split(play, ", ")
		for _, col := range set {
			item := strings.Split(col, " ")
			if fromBag, err := strconv.Atoi(item[0]); err == nil {
				if currentCubicle, ok := mapCube[string(item[1][0])]; ok && currentCubicle < fromBag {
					mapCube[string(item[1][0])] = fromBag
				}
			}
		}
		if !valid {
			break
		}
	}
	return valid, mapCube["r"] * mapCube["g"] * mapCube["b"]
}
