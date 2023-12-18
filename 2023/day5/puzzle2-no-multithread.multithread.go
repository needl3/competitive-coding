package day5

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/needl3/adventofcode/2023/utils"
)

func Puzzle2NoMultiThread() {
	file, _ := os.Open("day5/input.txt")

	// Preprocessing data
	line, _ := utils.Readline(file)
	var seeds = []int{}
	for _, sd := range strings.Split(strings.Split(line, ": ")[1], " ") {
		if n, err := strconv.Atoi(sd); err == nil {
			seeds = append(seeds, n)
		}
	}
	utils.Readline(file) // Ignore this extra spaceful line

	var stages []*Stage
	for {
		s, err := PrepareMap(file)
		stages = append(stages, s)
		if err != nil {
			break
		}
	}
	// Preprocess finished

	lowest := -1
	for i := 0; i < len(seeds); i += 2 {
		for seed := seeds[i]; seed < seeds[i]+seeds[i+1]; seed++ {
			// fmt.Printf("For seed: %d => ", seed)
			dst := seed
			for _, s := range stages {
				dst = s.Dest(dst)
				// fmt.Printf("%d => ", seed)
			}
			if lowest > dst || lowest < 0 {
				lowest = dst
			}
			// fmt.Println()
		}
	}

	fmt.Println("Final ans: ", lowest)
}
