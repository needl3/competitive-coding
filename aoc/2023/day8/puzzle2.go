package day8

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/needl3/adventofcode/2023/utils"
)

var mapStep map[rune]int = map[rune]int{
	'L': 0,
	'R': 1,
}

func Puzzle2() {
	file, _ := os.Open("day8/input.txt")

	steps, _ := utils.Readline(file)
	steps = strings.TrimSpace(steps)
	utils.Readline(file) // Ignore extra spaceful line

	var heads []string
	for {
		if line, err := utils.Readline(file); err == nil {
			data := strings.Split(line, " = ")
			state := data[0]
			children := strings.Split(data[1], ", ")
			lChild := children[0][1:]
			rChild := children[1][:len(children[1])-1]

			states[state] = [2]string{lChild, rChild}
			if state[2] == 'A' {
				heads = append(heads, state)
			}
		} else {
			break
		}
	}

	resultChan := make(chan int)
	var wg sync.WaitGroup
	for _, head := range heads {
		wg.Add(1)
		go func(hd string, ch chan<- int) {
			count := 0
			for ; ; count++ {
				if hd[2] == 'Z' {
					break
				} else {
					lr := count % len(steps)
					ind := mapStep[rune(steps[lr])]
					hd = states[hd][ind]
				}
			}
			resultChan <- count
		}(head, resultChan)
	}
	var countArr = []int{}
	mx := 0
	go func(w *sync.WaitGroup) {
		for c := range resultChan {
			countArr = append(countArr, c)
			if c > mx {
				mx = c
			}
			wg.Done()
		}
	}(&wg)

	wg.Wait()
	fmt.Println("Count: ", lcm(countArr, mx))
}

func lcm(f []int, mx int) int {
	lcm := mx
	div := false
	for !div {
		div = true
		for _, d := range f {
			if lcm%d != 0 {
				lcm += mx
				div = false
				break
			}
		}
	}
	return lcm
}
