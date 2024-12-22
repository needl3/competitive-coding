package main

import (
	"fmt"
	"os"

	"github.com/needl3/adventofcode/2024/day1"
)

var sols = []func(){
	day1.Puzzle1,
}

var puzzles = map[string]map[string]func(){}

func init() {
	counter := 0
	for i := 1; i <= 25 && counter < len(sols); i++ {
		puzzles[fmt.Sprintf("d%d", i)] = map[string]func(){}
		for j := 1; j <= 2 && counter < len(sols); j++ {
			puzzles[fmt.Sprintf("d%d", i)][fmt.Sprintf("p%d", j)] = sols[counter]
			counter++
		}
	}
}

func main() {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) == 2 {
		if fxn, ok := puzzles[argsWithoutProg[0]][argsWithoutProg[1]]; ok {
			fxn()
		} else {
			fmt.Println("[x] No solution found")
		}
	} else {
		fmt.Println("[x] Invalid command")
	}
}
