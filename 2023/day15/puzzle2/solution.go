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

	var box = map[int][]map[string]int{}
	if line, err := Readline(file); err == nil {
		splitted := strings.Split(line, ",")
		for _, inp := range splitted {
			label, i := findLabel(inp)
			operation := inp[i]
			focal := -1
			if operation == '=' {
				focal, _ = strconv.Atoi(string(inp[i+1]))
			}

			updateBox(&box, label, focal, operation)
		}
	}
	fmt.Println("Sum: ", findSum(&box))
}

func findLabel(inp string) (string, int) {
	i := 0
	label := ""
	for ; i < len(inp); i++ {
		if inp[i] == '=' || inp[i] == '-' {
			break
		}
		label += string(inp[i])
	}
	return label, i
}

func findSum(box *map[int][]map[string]int) int {
	sum := 0
	for k, v := range *box {
		for i, v2 := range v {
			for _, v := range v2 {
				sum += (k + 1) * (i + 1) * v
			}
		}
	}
	return sum
}

func updateBox(box *map[int][]map[string]int, label string, focal int, operation byte) {
	boxNo := findHash(label)
	_, ok := (*box)[boxNo]
	if !ok {
		(*box)[boxNo] = []map[string]int{}
	}

	switch operation {
	case '-':
		var replaced = []map[string]int{}
		for _, v := range (*box)[boxNo] {
			if _, ok := v[label]; !ok {
				replaced = append(replaced, v)
			}
		}
		(*box)[boxNo] = replaced
		break
	case '=':
		replaced := false
		for i, v := range (*box)[boxNo] {
			if _, ok := v[label]; ok {
				((*box)[boxNo])[i][label] = focal
				replaced = true
				break
			}
		}
		if !replaced {
			(*box)[boxNo] = append((*box)[boxNo], map[string]int{
				label: focal,
			})
		}
		break
	default:
		fmt.Println("Invalid operation")
	}
}

func findHash(inp string) int {
	hs := 0
	for _, c := range inp {
		hs = ((hs + int(c)) * 17) % 256
	}
	return hs
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
