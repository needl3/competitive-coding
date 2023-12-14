package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"sync"
)

var mapStep map[rune]int = map[rune]int{
	'L': 0,
	'R': 1,
}

var states = map[string][2]string{}

func main() {
	// file, _ := os.Open("../input.txt")
	file, _ := os.Open("../input-submit.txt")

	steps, _ := Readline(file)
	steps = strings.TrimSpace(steps)
	Readline(file) // Ignore extra spaceful line

	var heads []string
	for {
		if line, err := Readline(file); err == nil {
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
	fmt.Println(countArr)
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
