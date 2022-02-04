package configparser

import "RainbowRunner/internal/database"

type DRNamespaceStack struct {
	Stack []*database.DRNamespace
	Index int
}

func NewDRNamespaceStack() *DRNamespaceStack {
	return &DRNamespaceStack{
		Stack: make([]*database.DRNamespace, 1024),
		Index: 0,
	}
}

func (s *DRNamespaceStack) Push(str *database.DRNamespace) {
	s.Stack[s.Index] = str
	s.Index = s.Index + 1
}

func (s *DRNamespaceStack) Pop() *database.DRNamespace {
	s.Index = s.Index - 1
	res := s.Stack[s.Index]
	return res
}

func (s *DRNamespaceStack) Current() *database.DRNamespace {
	if s.Index == 0 {
		return nil
	}

	return s.Stack[s.Index-1]
}

func (s *DRNamespaceStack) SetCurrent(i *database.DRNamespace) {
	s.Stack[s.Index-1] = i
}

func (s *DRNamespaceStack) Previous() *database.DRNamespace {
	if len(s.Stack) >= 2 {
		return s.Stack[s.Index-2]
	}

	return nil
}
