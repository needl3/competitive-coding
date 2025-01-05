package day3

import "fmt"

func Puzzle2() {
	input := ParseInput("day3/input.txt")
	fsm := NewFSM()
	fsm2 := NewFSMV2(&fsm)
	for _, x := range input {
		fsm2.read(x)
	}
	fmt.Println("Answer: ", fsm.CumValue)
}

type FSMV2 struct {
	fsm          *FSM
	StateMap     map[rune]int8
	CurrentState int8
	EnableMode   bool
}

func NewFSMV2(fsm *FSM) FSMV2 {
	return FSMV2{
		fsm: fsm,
		StateMap: map[rune]int8{
			'd':  0,
			'o':  1,
			'n':  2,
			'\'': 3,
			't':  4,
			'(':  5,
			')':  6,
		},
	}
}

func (f *FSMV2) reset() {
	f.CurrentState = 0
}

func (f *FSMV2) finalize() {
	if f.EnableMode {
		f.fsm.EnableMul()
	} else {
		f.fsm.DisableMul()
	}
	f.reset()
}

func (f *FSMV2) advance(x rune) {

	if f.CurrentState == 6 {
		f.finalize()
	} else {
		if x == '(' && f.CurrentState == 2 {
			f.CurrentState = 6
			f.EnableMode = true
		} else {
			if x == '(' {
				f.EnableMode = false
			}
			f.CurrentState++
		}
	}
}

func (f *FSMV2) read(x rune) {
	state, ok := f.StateMap[x]
	if ok {
		if f.CurrentState == state || (x == '(' && f.CurrentState == 2) {
			f.advance(x)
		}
	} else {
		f.reset()
	}
	f.fsm.read(x)
}
