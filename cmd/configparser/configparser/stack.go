package configparser

import (
	"RainbowRunner/internal/types/configtypes"
)

type DRClassStack struct {
	Stack []*configtypes.DRClass
	Index int
}

func NewDRClassStack() *DRClassStack {
	return &DRClassStack{
		Stack: make([]*configtypes.DRClass, 1024),
		Index: 0,
	}
}

func (s *DRClassStack) Push(str *configtypes.DRClass) {
	s.Stack[s.Index] = str
	s.Index = s.Index + 1
}

func (s *DRClassStack) Pop() *configtypes.DRClass {
	s.Index = s.Index - 1
	res := s.Stack[s.Index]
	return res
}

func (s *DRClassStack) Current() *configtypes.DRClass {
	if s.Index == 0 {
		return nil
	}

	return s.Stack[s.Index-1]
}

func (s *DRClassStack) SetCurrent(i *configtypes.DRClass) {
	s.Stack[s.Index-1] = i
}

func (s *DRClassStack) Previous() *configtypes.DRClass {
	if len(s.Stack) >= 2 {
		return s.Stack[s.Index-2]
	}

	return nil
}
