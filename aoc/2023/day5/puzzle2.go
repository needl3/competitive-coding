package day5

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/needl3/adventofcode/2023/utils"
)

func Puzzle2() {
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
