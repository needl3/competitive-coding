package day3

import (
	"log"
	"os"

	"github.com/needl3/adventofcode/2024/utils"
)

func ParseInput(filePath string) string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Panicf("Couldn't read input")
	}

	s := ""
	for {
		line, err := utils.Readline(file)
		if err != nil {
			break
		}
		s += line
	}
	return s
}
