package day6

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/needl3/adventofcode/2023/utils"
)

type Factors struct {
	P1 int
	P2 int
}

func (f Factors) prod() int {
	return f.P1 * f.P2
}
func proceed(f *Factors) {
	f.P1 -= 1
	f.P2 += 1
}

func Puzzle1() {
	file, _ := os.Open("day6/input.txt")

	str, _ := utils.Readline(file)
	t := filterNumbers(str)

	str, _ = utils.Readline(file)
	hS := filterNumbers(str)

	ways := -1
	for i := range t {
		maxProd := findMaxProducts(t[i])
		fact := findFactors(t[i], hS[i])
		if ways < 0 {
			ways = findWays(maxProd, fact, hS[i])
		} else {
			ways *= findWays(maxProd, fact, hS[i])
		}
	}
	fmt.Println("Final ways: ", ways)

}

func findWays(myFact *Factors, highScoreFact *Factors, highScore int) int {
	count := 0
	if highScoreFact.prod() == highScore {
		proceed(highScoreFact)
	}
	for highScoreFact.prod() > highScore {
		proceed(highScoreFact)
		count++
	}

	return count
}

func findFactors(time int, score int) *Factors {
	i := 1
	j := time - i
	for i*j < score {
		i++
		j = time - i
	}
	if i > j {
		return &Factors{
			P1: int(i),
			P2: int(j),
		}
	}
	return &Factors{
		P1: int(j),
		P2: int(i),
	}
}

func findMaxProducts(i int) *Factors {
	return &Factors{
		P1: int(math.Ceil(float64(i) / 2)),
		P2: int(math.Floor(float64(i) / 2)),
	}
}

func filterNumbers(str string) []int {
	var t []int
	splitted := strings.Split(str, ": ")[1]
	trimmed := strings.TrimSpace(splitted)
	trimmedSplitted := strings.Split(trimmed, " ")
	for _, tStr := range trimmedSplitted {
		if n, err := strconv.Atoi(tStr); err == nil {
			t = append(t, n)
		}
	}
	return t
}
