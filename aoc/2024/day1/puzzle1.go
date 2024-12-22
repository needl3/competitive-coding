package day1

import (
	"log"
	"slices"
)

func Puzzle1() {
	parsedData := ParseInput("day1/input.txt")

	slices.SortFunc(parsedData.LeftArr, func(a, b int) int {
		return a - b
	})

	slices.SortFunc(parsedData.RightArr, func(a, b int) int {
		return a - b
	})

	sum := 0
	for i, elem := range parsedData.LeftArr {
		sum += findAbsoluteDiff(elem, parsedData.RightArr[i])
	}
	log.Println("Answer: ", sum)
}

func findAbsoluteDiff(a, b int) int {
	if a < b {
		return b - a
	} else {
		return a - b
	}
}
