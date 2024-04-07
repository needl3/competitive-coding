package anagram

import (
	"fmt"
)

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	arr := make([]int, 26)
	for i, r := range s {
		arr[r-'a']++
		arr[t[i]-'a']--
	}

	finalRes := 0
	for _, k := range arr {
		finalRes |= k
	}
	return finalRes == 0
}

// Use this function to parse, prepare and format input from stdin
func Exec() {
	fmt.Println(isAnagram("gramana", "anagram"))
}
