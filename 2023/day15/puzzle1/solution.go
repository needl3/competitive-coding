package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func main() {
	// file, _ := os.Open("../input.txt")
	file, _ := os.Open("../input-submit.txt")

	gSum := 0
	if line, err := Readline(file); err == nil {
		splitted := strings.Split(line, ",")
		for _, inp := range splitted {
			gSum += findHash(inp)
		}
	}
	fmt.Println("Hash: ", gSum)
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
