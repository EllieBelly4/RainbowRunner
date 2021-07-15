package datatypes

type StringStack struct {
	Stack []string
	Index int
}

func NewStringStack() *StringStack {
	return &StringStack{
		Stack: make([]string, 1024),
		Index: 0,
	}
}

func (s *StringStack) Push(str string) {
	s.Stack[s.Index] = str
	s.Index = s.Index + 1
}

func (s *StringStack) Pop() string {
	s.Index = s.Index - 1
	res := s.Stack[s.Index]
	return res
}

type IntStack struct {
	Stack []int
	Index int
}

func NewIntStack() *IntStack {
	return &IntStack{
		Stack: make([]int, 1024),
		Index: 0,
	}
}

func (s *IntStack) Push(str int) {
	s.Stack[s.Index] = str
	s.Index = s.Index + 1
}

func (s *IntStack) Pop() int {
	s.Index = s.Index - 1
	res := s.Stack[s.Index]
	return res
}

func (s *IntStack) Current() int {
	return s.Stack[s.Index-1]
}

func (s *IntStack) SetCurrent(i int) {
	s.Stack[s.Index-1] = i
}
