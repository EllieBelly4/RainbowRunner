package configparser

import (
	"RainbowRunner/internal/types/drconfigtypes"
)

type DRClassStack struct {
	Stack []*drconfigtypes.DRClass
	Index int
}

func NewDRClassStack() *DRClassStack {
	return &DRClassStack{
		Stack: make([]*drconfigtypes.DRClass, 1024),
		Index: 0,
	}
}

func (s *DRClassStack) Push(str *drconfigtypes.DRClass) {
	s.Stack[s.Index] = str
	s.Index = s.Index + 1
}

func (s *DRClassStack) Pop() *drconfigtypes.DRClass {
	s.Index = s.Index - 1
	res := s.Stack[s.Index]
	return res
}

func (s *DRClassStack) Current() *drconfigtypes.DRClass {
	if s.Index == 0 {
		return nil
	}

	return s.Stack[s.Index-1]
}

func (s *DRClassStack) SetCurrent(i *drconfigtypes.DRClass) {
	s.Stack[s.Index-1] = i
}

func (s *DRClassStack) Previous() *drconfigtypes.DRClass {
	if len(s.Stack) >= 2 {
		return s.Stack[s.Index-2]
	}

	return nil
}
