package day1

import "log"

func Puzzle2() {
	parsedData := ParseInput("day1/input.txt")

	rightElementCountMap := countSimilarNumbers(parsedData.RightArr)
	sum := 0
	for _, n := range parsedData.LeftArr {
		sum += n * rightElementCountMap[n]
	}
	log.Println("Answer: ", sum)
}

func countSimilarNumbers(nums []int) map[int]int {
	var count = map[int]int{}
	for _, n := range nums {
		if _, ok := count[n]; ok {
			count[n]++
		} else {
			count[n] = 1
		}

	}
	return count
}
