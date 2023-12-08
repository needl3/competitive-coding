package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func PrepareMap(f *os.File) (*Stage, error) {
	_, _ = Readline(f) // Ignore this line specifying title
	// fmt.Println(title)

	stg := Stage{}
	for {
		if str, err := Readline(f); err == nil {
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

func main() {
	// file, _ := os.Open("../input.txt")
	file, _ := os.Open("../input-submit.txt")

	// Preprocessing data
	line, _ := Readline(file)
	var seeds = []int{}
	for _, sd := range strings.Split(strings.Split(line, ": ")[1], " ") {
		if n, err := strconv.Atoi(sd); err == nil {
			seeds = append(seeds, n)
		}
	}
	Readline(file) // Ignore this extra spaceful line

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