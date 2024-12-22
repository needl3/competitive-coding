package day1

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/needl3/adventofcode/2024/utils"
)

type ParsedData struct {
	LeftArr  []int
	RightArr []int
}

func ParseInput(filePath string) ParsedData {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to open file")
	}

	leftArray := []int{}
	rightArray := []int{}

	for {
		line, err := utils.Readline(file)
		if err != nil {
			break
		}
		splittedLine := strings.Split(line, "   ")
		leftInt, err := strconv.Atoi(splittedLine[0])
		rightInt, err := strconv.Atoi(splittedLine[1])
		if err != nil {
			log.Fatal("Invlid number on input file")
		}
		leftArray = append(leftArray, leftInt)
		rightArray = append(rightArray, rightInt)
	}
	return ParsedData{
		LeftArr:  leftArray,
		RightArr: rightArray,
	}
}
