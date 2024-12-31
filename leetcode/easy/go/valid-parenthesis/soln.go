package validparenthesis

import (
	"errors"
)

type Stack struct {
	size int
	data []rune
	top  int
}

func (s *Stack) Empty() bool {
	return s.top == 0
}

func (s *Stack) Push(x rune) error {
	if s.top == s.size {
		return errors.New("Cannot push. Stack is full")
	}
	s.data[s.top] = x
	s.top++
	return nil
}

func (s *Stack) Pop() (rune, error) {
	if s.top == 0 {
		return 'x', errors.New("Cannot pop. Stack is empty")
	}
	s.top--
	return s.data[s.top], nil
}

var m = map[rune]rune{')': '(', '}': '{', ']': '['}

func isValid(s string) bool {
	if len(s)%2 == 1 {
		return false
	}

	stack := Stack{
		size: len(s),
		data: make([]rune, len(s), len(s)),
	}

	for _, x := range s {
		switch x {
		case '(', '[', '{':
			if err := stack.Push(x); err != nil {
				return false
			}
		case ']', ')', '}':
			opener := m[x]
			if v, err := stack.Pop(); err != nil || v != opener {
				return false
			}
		}
	}
	return stack.Empty()
}
