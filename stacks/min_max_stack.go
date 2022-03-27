package stacks

// MinMaxStack
//  Write a 'MinMaxStack' class for a Min Max Stack. The class should
//  support:
//
//  Pushing and popping values on and off the stack.
//  Peeking at the value at the top of the stack.
//
//    Getting both the minimum and the maximum values in the stack at any given
//    point in time.
// Stack is a LiFo, Last in, First out
type MinMaxStack struct {
	stack  []int
	minMax []entry
}

type entry struct {
	min int
	max int
}

func NewMinMaxStack() *MinMaxStack {
	return &MinMaxStack{}
}

func (s *MinMaxStack) Peek() int {
	return s.stack[len(s.stack)-1]
}

func (s *MinMaxStack) Pop() int {
	val := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack) - 1]
	s.minMax = s.minMax[:len(s.minMax) - 1]
	return val
}

func (s *MinMaxStack) Push(number int) {
	e := entry{
		min: number,
		max: number,
	}

	if len(s.stack)>0{
		e.min=s.min(e.min, s.minMax[len(s.stack)-1].min)
		e.max=s.max(e.max, s.minMax[len(s.stack)-1].max)
	}
	s.stack = append(s.stack, number)
	s.minMax = append(s.minMax, e)
}

func (s *MinMaxStack) GetMin() int {
	return s.minMax[len(s.stack)-1].min
}

func (s *MinMaxStack) GetMax() int {
	return s.minMax[len(s.stack)-1].max
}

func (s *MinMaxStack) max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func (s *MinMaxStack) min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
