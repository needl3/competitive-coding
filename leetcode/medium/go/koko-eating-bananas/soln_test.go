package kokoeatingbananas

import (
	"testing"
)

func TestSoln0(t *testing.T) {
	valid := isValid(5, 3, []int{2, 5, 3})
	if !valid {
		t.Fail()
	}
}

func TestSoln(t *testing.T) {
	speed := minEatingSpeed([]int{3, 6, 7, 11}, 8)
	if speed != 4 {
		t.Log("First test: ", speed)
		t.Fail()
	}
}

func TestSoln2(t *testing.T) {
	speed := minEatingSpeed([]int{30, 11, 23, 4, 20}, 6)
	if speed != 23 {
		t.Log("Second test: ", speed)
		t.Fail()
	}
}

func TestSoln3(t *testing.T) {
	speed := minEatingSpeed([]int{3, 3, 3, 3}, 5)
	if speed != 3 {
		t.Log("Third test: ", speed)
		t.Fail()
	}
}

func TestSoln4(t *testing.T) {
	speed := minEatingSpeed([]int{312884470}, 312884469)
	t.Log(speed)
	if speed != 2 {
		t.Log("Fourth test: ", speed)
		t.Fail()
	}
}

func TestSoln5(t *testing.T) {
	speed := minEatingSpeed([]int{332484035, 524908576, 855865114, 632922376, 222257295, 690155293, 112677673, 679580077, 337406589, 290818316, 877337160, 901728858, 679284947, 688210097, 692137887, 718203285, 629455728, 941802184}, 312884469)
	t.Log(speed)
	if speed != 35 {
		t.Fail()
	}
}
