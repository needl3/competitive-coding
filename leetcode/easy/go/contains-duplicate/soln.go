package containsduplicate

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func containsDuplicate(nums []int) bool {
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i]^nums[i+1] == 0 {
			return true
		}
	}
	return false
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
	fmt.Println(containsDuplicate(formattedArr))
}
