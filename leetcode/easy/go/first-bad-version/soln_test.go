package firstbadversion

import "testing"

func TestSoln(t *testing.T) {
	if firstBadVersion(100) != 1 {
		t.Fail()
	}
}
