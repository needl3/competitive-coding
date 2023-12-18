package day4

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/needl3/adventofcode/2023/utils"
)

func Puzzle1() {
	file, _ := os.Open("day4/input.txt")

	sum := 0
	for {
		wonMap := map[int]bool{}
		if line, err := utils.Readline(file); err == nil {
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
