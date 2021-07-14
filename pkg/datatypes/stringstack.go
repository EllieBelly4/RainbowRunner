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
