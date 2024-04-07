package groupanagrams

import "fmt"

var PRIMES = []int{5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107}

func groupAnagrams(strs []string) [][]string {
	m := map[int][]string{}
	for _, str := range strs {
		hash := gimmeHash(str)
		if _, ok := m[hash]; ok {
			m[hash] = append(m[hash], str)
		} else {
			m[hash] = []string{str}
		}
	}
	result := [][]string{}
	for _, e := range m {
		result = append(result, e)
	}
	return result
}

// This should generate same hash for anagrams
// Uses prime multiplication for unique hash but same of anagrams
// This is limited to small strings only
// As in question s[i].length < 100 so it's good
func gimmeHash(s string) int {
	mul := 1
	for _, s := range s {
		mul *= PRIMES[s-'a']
	}
	return mul 
}

func Exec() {
	fmt.Println(groupAnagrams([]string{"cab", "tin", "pew", "duh", "may", "ill", "buy", "bar", "max", "doc"}))
}
