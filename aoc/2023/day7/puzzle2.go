package day7

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var ans string = "251857682H, 251930388H, 251692783X, 252056944H, 251740281X"

func Puzzle2() {
	file, _ := os.Open("day7/input.txt")

	game := map[string]int{}
	var cards []ParsedCardParentKind
	for {
		if line, err := Readline(file); err == nil {
			splitted := strings.Split(line, " ")
			game[splitted[0]], _ = strconv.Atoi(splitted[1])
			maxHand := getMaxHand(splitted[0])
			cards = append(cards, ParsedCardParentKind{
				Orig: ParsedCardKind{
					Name:     splitted[0],
					Strength: parseKindOneStrength(splitted[0]),
				},
				Maxed: ParsedCardKind{
					Name:     maxHand,
					Strength: parseKindOneStrength(maxHand),
				},
			})
		} else {
			break
		}
	}

	sortV2(cards)

	sum := 0
	for i, card := range cards {
		sum += game[card.Orig.Name] * (i + 1)
	}

	fmt.Println("Sum: ", sum)
}

func getMaxHand(hand string) string {
	a := map[rune]int{}
	for _, letter := range hand {
		a[letter]++
	}
	maxCount := 0
	maxChar := 'J'
	for l, i := range a {
		if i > maxCount && l != 'J' {
			maxCount = i
			maxChar = l
		}
	}
	x := strings.ReplaceAll(hand, "J", string(maxChar))
	return x
}

type ParsedCardParentKind struct {
	Orig  ParsedCardKind
	Maxed ParsedCardKind
}

func sortV2(cards []ParsedCardParentKind) {
	// Do a simple bubble sort for foggs sake
	for i := 0; i < len(cards)-1; i++ {
		swapped := false
		for j := 0; j < len(cards)-1; j++ {
			if swp := sortByFirstOrderingV2(cards[j], cards[j+1]); swp {
				swapped = swp
				tmp := cards[j]
				cards[j] = cards[j+1]
				cards[j+1] = tmp
			}
		}
		if !swapped {
		}
	}
}

func sortByFirstOrderingV2(c1 ParsedCardParentKind, c2 ParsedCardParentKind) bool {
	if c1.Maxed.Strength > c2.Maxed.Strength {
		return true
	} else if c1.Maxed.Strength == c2.Maxed.Strength {
		return sortBySecondOrderingV2(c1.Orig, c2.Orig)
	}
	return false
}

var cardStrengthMapV2 map[rune]int = map[rune]int{
	'A': 12, 'K': 11, 'Q': 10, 'J': -1, 'T': 8,
	'9': 7, '8': 6, '7': 5, '6': 4, '5': 3, '4': 2, '3': 1, '2': 0,
}

func sortBySecondOrderingV2(str1 ParsedCardKind, str2 ParsedCardKind) bool {
	for i := range str1.Name {
		n1, _ := cardStrengthMapV2[rune(str1.Name[i])]
		n2, _ := cardStrengthMapV2[rune(str2.Name[i])]
		if n1 > n2 {
			return true
		} else if n1 < n2 {
			return false
		}
	}
	return false
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
