package climbingstairs

import (
	"testing"
)

func TestSoln(t *testing.T) {
	steps := climbStairs(3)
	if steps != 3 {
		t.Fail()
	}
}
