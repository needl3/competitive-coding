package main

import (
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
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

func main() {
	// file, _ := os.Open("../input.txt")
	file, _ := os.Open("../input-submit.txt")

	str, _ := Readline(file)
	t := filterNumbers(str)

	str, _ = Readline(file)
	hS := filterNumbers(str)

	ways := -1
	for i, _ := range t {
		maxProd := findMaxProducts(t[i])
		fact := findFactors(t[i], hS[i])
		// fmt.Println("Score prod: ", fact.P1, fact.P2, "=", fact.P1*fact.P2)
		// fmt.Println("My prod: ", maxProd.P1, maxProd.P2, "=", maxProd.P1*maxProd.P2)
		// fmt.Println("High score: ", hS[i])
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
	// fmt.Println("Local ways: ", count)

	return count
}

func findWaysV1(myFact *Factors, highScoreFact *Factors, highScore int) int {
	// Wrong approach
	if highScoreFact.prod() == highScore {
		return (myFact.P1-highScoreFact.P2-1)*2 + 1
	} else {
		return (myFact.P1 - highScoreFact.P2) * 2
	}
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
