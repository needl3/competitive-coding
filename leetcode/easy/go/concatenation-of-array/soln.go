package singlenumber

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getConcatenation(nums []int) []int {
	finalNums := make([]int, len(nums)*2)
	fmt.Println(nums)
	for i := 0; i < len(nums)*2; i++ {
		finalNums[i] = nums[i%len(nums)]
	}
	return finalNums
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

	// Call the function provided from leetcode to invoke solution function
	fmt.Println(getConcatenation(formattedArr))
}
