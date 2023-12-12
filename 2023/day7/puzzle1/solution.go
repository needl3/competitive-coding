package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// file, _ := os.Open("../input.txt")
	file, _ := os.Open("../input-submit.txt")

	game := map[string]int{}
	var cards []ParsedCardKind
	for {
		if line, err := Readline(file); err == nil {
			splitted := strings.Split(line, " ")
			game[splitted[0]], _ = strconv.Atoi(splitted[1])
			cards = append(cards, ParsedCardKind{
				Name:     splitted[0],
				Strength: parseKindOneStrength(splitted[0]),
			})
		} else {
			break
		}
	}

	sort(cards)

	sum := 0
	for i, card := range cards {
		sum += game[card.Name] * (i + 1)
	}

	fmt.Println("Sum: ", sum)
}

type ParsedCardKind struct {
	Name     string
	Strength int
}

func parseKindOneStrength(str string) int {
	a := map[rune]int{}
	for _, letter := range str {
		a[letter]++
	}

	if len(a) == len(str) {
		return 1
	} else if len(a) == len(str)-1 {
		return 2
	} else if len(a) == 1 {
		return 7
	} else if len(a) == 2 {
		for _, v := range a {
			if v == 4 || v == 1 {
				return 6
			} else {
				return 5
			}
		}
		return 3 // This won't happen acc to input
	} else {
		for _, v := range a {
			if v == 2 {
				return 3
			}
		}
		return 4
	}
}

func sort(cards []ParsedCardKind) {
	// Do a simple bubble sort for foggs sake
	for i := 0; i < len(cards)-1; i++ {
		swapped := false
		for j := 0; j < len(cards)-1; j++ {
			swp := swapped
			cards[j], cards[j+1], swp = sortByFirstOrdering(cards[j], cards[j+1])
			if swp {
				swapped = swp
			}
		}
		if !swapped {
			break
		}
	}
}

func sortByFirstOrdering(c1 ParsedCardKind, c2 ParsedCardKind) (ParsedCardKind, ParsedCardKind, bool) {
	if c1.Strength > c2.Strength {
		return c2, c1, true
	} else if c1.Strength == c2.Strength {
		return sortBySecondOrdering(c1, c2)
	}
	return c1, c2, false
}

var cardStrengthMap map[rune]int = map[rune]int{
	'A': 12, 'K': 11, 'Q': 10, 'J': 9, 'T': 8,
	'9': 7, '8': 6, '7': 5, '6': 4, '5': 3, '4': 2, '3': 1, '2': 0,
}

func sortBySecondOrdering(str1 ParsedCardKind, str2 ParsedCardKind) (ParsedCardKind, ParsedCardKind, bool) {
	for i := range str1.Name {
		n1, _ := cardStrengthMap[rune(str1.Name[i])]
		n2, _ := cardStrengthMap[rune(str2.Name[i])]
		if n1 > n2 {
			return str2, str1, true
		} else if n1 < n2 {
			return str1, str2, false
		}
	}
	return str1, str2, false
}

func parseBid(subSplitted []string) int {
	num := 0
	for _, nStr := range subSplitted {
		for _, i := range nStr {
			if n, err := strconv.Atoi(string(i)); err == nil {
				num = num*10 + n
			}

		}
	}
	return num
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
