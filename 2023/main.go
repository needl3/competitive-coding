package main

import (
	"fmt"
	"os"

	"github.com/needl3/adventofcode/2023/day1"
	"github.com/needl3/adventofcode/2023/day10"
	"github.com/needl3/adventofcode/2023/day14"
	"github.com/needl3/adventofcode/2023/day15"
	"github.com/needl3/adventofcode/2023/day16"
	"github.com/needl3/adventofcode/2023/day18"
	"github.com/needl3/adventofcode/2023/day2"
	"github.com/needl3/adventofcode/2023/day3"
	"github.com/needl3/adventofcode/2023/day4"
	"github.com/needl3/adventofcode/2023/day5"
	"github.com/needl3/adventofcode/2023/day6"
	"github.com/needl3/adventofcode/2023/day7"
	"github.com/needl3/adventofcode/2023/day8"
	"github.com/needl3/adventofcode/2023/day9"
)

var sols = []func(){
	day1.Puzzle1, day1.Puzzle2,
	day2.Puzzle1, day2.Puzzle2,
	day3.Puzzle1, day3.Puzzle2,
	day4.Puzzle1, day4.Puzzle2,
	day5.Puzzle1, day5.Puzzle2,
	day6.Puzzle1, day6.Puzzle2,
	day7.Puzzle1, day7.Puzzle2,
	day8.Puzzle1, day8.Puzzle2,
	day9.Puzzle1, day9.Puzzle2,
	day10.Puzzle1, day10.Puzzle2,
	func() {}, func() {},
	func() {}, func() {},
	func() {}, func() {},
	day14.Puzzle1, day14.Puzzle2,
	day15.Puzzle1, day15.Puzzle2,
	day16.Puzzle1, day16.Puzzle2,
	func() {}, func() {},
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
	argsWithoutProg := os.Args[2:]
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
