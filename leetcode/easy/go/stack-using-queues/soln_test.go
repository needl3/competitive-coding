package stackusingqueues

import (
	"fmt"
	"testing"
)

func TestSoln(t *testing.T) {
	s := Constructor()
	s.Push(1)
	fmt.Println(s.Top())
	s.Push(2)
	fmt.Println(s.Top())
	s.Push(3)
	fmt.Println(s.Top())
	s.Pop()
	fmt.Println(s.Top())
	s.Push(4)
	fmt.Println(s.Top())
}
