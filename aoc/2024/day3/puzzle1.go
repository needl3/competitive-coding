package day3

import (
	"fmt"
	"strconv"
)

func Puzzle1() {
	input := ParseInput("day3/input.txt")
	fsm := NewFSM()
	for _, x := range input {
		fsm.read(x)
	}
	fmt.Println("Answer: ", fsm.CumValue)
}

type TStateMap map[rune]int8

func NewFSM() FSM {
	return FSM{
		MulEnabled: true,
		StateMap: TStateMap{
			'm': 0,
			'u': 1,
			'l': 2,
			'(': 3,
		},
	}
}

func (f *FSM) EnableMul() {
	f.MulEnabled = true
}

func (f *FSM) DisableMul() {
	f.MulEnabled = false
}

type FSM struct {
	MulEnabled   bool
	CurrentState int8
	CumValue     int
	MulV1        int
	MulV2        int
	StateMap     TStateMap
}

func (f *FSM) reset(v rune) {
	f.CurrentState = 0
	f.MulV1 = 0
	f.MulV2 = 0
	if v == 'm' {
		f.CurrentState++
	}
}

func (f *FSM) advance() {
	f.CurrentState++
}

func (f *FSM) checkIntegerState(x rune) {
	if f.CurrentState == 4 {
		val, err := strconv.Atoi(string(x))
		if err != nil {
			f.reset(x)
		} else {
			f.MulV1 = 10*f.MulV1 + val
		}
		return
	} else if f.CurrentState == 6 {
		val, err := strconv.Atoi(string(x))
		if err != nil {
			f.reset(x)
		} else {
			f.MulV2 = 10*f.MulV2 + val
		}
		return
	}
	f.reset(x)
}

func (f *FSM) finalize() {
	if f.MulEnabled {
		f.CumValue += f.MulV1 * f.MulV2
	}
	f.reset('@')
}

func (f *FSM) read(x rune) {
	switch {
	case x == ',' && f.CurrentState == 4:
		f.advance()
		f.advance()
	case x == ')' && f.CurrentState == 6:
		f.finalize()
	default:
		if f.CurrentState == 4 || f.CurrentState == 6 {
			f.checkIntegerState(x)
			return
		}

		if v, ok := f.StateMap[x]; ok {
			if f.CurrentState == v {
				f.advance()
			}
		} else {
			f.reset('@')
		}
	}
}
