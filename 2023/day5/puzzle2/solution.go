package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
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

	gLowest := -1
	var mx sync.Mutex
	resultChan := make(chan int)
	var wg sync.WaitGroup

	for i := 0; i < len(seeds)-2; i += 2 {
		wg.Add(1)
		go func(j int, resChan chan<- int) {

			fmt.Printf("From seed: %d => %d\n", seeds[j], seeds[j]+seeds[j+1]-1)

			lowest := -1
			for seed := seeds[j]; seed < seeds[j]+seeds[j+1]; seed++ {
				dst := seed
				for _, s := range stages {
					dst = s.Dest(dst)
				}
				if lowest > dst || lowest < 0 {
					lowest = dst
				}
			}

			resChan <- lowest
		}(i, resultChan)
	}

	go func() {
		for {
			result := <-resultChan
			fmt.Println("Received local lowest result: ", result)
			mx.Lock()
			if gLowest > result || gLowest < 0 {
				gLowest = result
			}
			mx.Unlock()
			wg.Done()
		}
	}()
	wg.Wait()

	fmt.Println("Final global lowest: ", gLowest)
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
