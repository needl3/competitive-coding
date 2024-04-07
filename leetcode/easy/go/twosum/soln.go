package twosum

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func twoSum(nums []int, target int) []int {
	whatINeed := map[int]int{}
	for i := 1; i < len(nums); i++ {
		whatINeed[target-nums[i-1]] = i - 1
		if _, ok := whatINeed[nums[i]]; ok {
			return []int{whatINeed[nums[i]], i}
		}
	}
	return []int{-1, -1}
}

// Use this function to parse, prepare and format input from stdin
func Exec() {
	reader := bufio.NewScanner(os.Stdin)
	reader.Scan()
	inp := reader.Text()
	arr := strings.Split(inp[1:], ",")
	var formattedArr []int
	for _, i := range arr {
		formattedArr = append(formattedArr, int(i[0]-'0'))
	}
	reader.Scan()

	// Call the function provided from leetcode to invoke solution function
	fmt.Println(twoSum(formattedArr, int(reader.Text()[0]-'0')))
}
