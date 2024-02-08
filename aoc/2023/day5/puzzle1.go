package day5

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/needl3/adventofcode/2023/utils"
)

type StageType interface {
	Dest(int) int
}

type Std struct {
	Source      int
	Destination int
	Range       int
}

type Stage struct {
	maps []Std
}

func (stg Stage) Dest(source int) int {
	for _, s := range stg.maps {
		if source >= s.Source && source < s.Source+s.Range {
			return s.Destination + (source - s.Source)
		}
	}
	return source
}

func Puzzle1() {
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
	for _, seed := range seeds {
		// fmt.Printf("For seed: %d => ", seed)
		for _, s := range stages {
			seed = s.Dest(seed)
			// fmt.Printf("%d => ", seed)
		}
		if lowest > seed || lowest < 0 {
			lowest = seed
		}
		// fmt.Println()
	}
	fmt.Println("Final ans: ", lowest)
}

func PrepareMap(f *os.File) (*Stage, error) {
	_, _ = utils.Readline(f) // Ignore this line specifying title
	// fmt.Println(title)

	stg := Stage{}
	for {
		if str, err := utils.Readline(f); err == nil {
			if len(str) > 0 {
				splitted := strings.Split(strings.TrimSpace(str), " ")
				source, _ := strconv.Atoi(splitted[1])
				destination, _ := strconv.Atoi(splitted[0])
				rng, _ := strconv.Atoi(splitted[2])
				stg.maps = append(stg.maps, Std{
					Source:      source,
					Destination: destination,
					Range:       rng,
				})
			} else {
				break
			}
		} else {
			return &stg, errors.New("EOF")
		}
	}
	// fmt.Println(stg)

	return &stg, nil
}
