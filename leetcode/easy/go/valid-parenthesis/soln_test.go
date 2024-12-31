package validparenthesis

import "testing"

func TestSoln1(t *testing.T) {
	r := isValid("({})")
	if !r {
		t.FailNow()
	}
}

func TestSoln2(t *testing.T) {
	r := isValid("({[[]})")
	if r {
		t.FailNow()
	}
}

func TestSoln3(t *testing.T) {
	r := isValid("(")
	if r {
		t.FailNow()
	}
}

func TestSoln4(t *testing.T) {
	r := isValid("")
	if !r {
		t.FailNow()
	}
}

func TestSoln5(t *testing.T) {
	r := isValid("()))))")
	if r {
		t.FailNow()
	}
}
