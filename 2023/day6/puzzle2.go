package day6

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/needl3/adventofcode/2023/utils"
)

func Puzzle2() {
	file, _ := os.Open("day6/input.txt")

	str, _ := utils.Readline(file)
	t := filterNumbersV2(str)

	str, _ = utils.Readline(file)
	hS := filterNumbersV2(str)

	ways := -1
	maxProd := findMaxProducts(t)
	fact := findFactors(t, hS)
	if ways < 0 {
		ways = findWays(maxProd, fact, hS)
	} else {
		ways *= findWays(maxProd, fact, hS)
	}
	fmt.Println("Final ways: ", ways)

}

func filterNumbersV2(str string) int {
	t := 0
	splitted := strings.Split(str, ": ")[1]
	for _, tStr := range splitted {
		if n, err := strconv.Atoi(string(tStr)); err == nil {
			t = 10*t + n
		}
	}
	return t
}
