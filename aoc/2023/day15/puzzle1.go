package day15

import (
	"fmt"
	"os"
	"strings"

	"github.com/needl3/adventofcode/2023/utils"
)

func Puzzle1() {
	file, _ := os.Open("day15/input.txt")

	gSum := 0
	if line, err := utils.Readline(file); err == nil {
		splitted := strings.Split(line, ",")
		for _, inp := range splitted {
			gSum += findHash(inp)
		}
	}
	fmt.Println("Hash: ", gSum)
}

func findHash(inp string) int {
	hs := 0
	for _, c := range inp {
		hs = ((hs + int(c)) * 17) % 256
	}
	return hs
}
