package integer

import "errors"

// Stack 栈
type Stack struct {
	size int
	data []int
}

// NewStack 新建栈
func NewStack(data ...int) *Stack {
	return &Stack{size: len(data), data: data}
}

// IsEmpty 是否为空栈
func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

// Size 大小
func (s *Stack) Size() int {
	return s.size
}

// Push 入栈
func (s *Stack) Push(data ...int) {
	if s.IsEmpty() {
		s.data = []int{}
	}
	realSize := len(s.data)
	if s.size == realSize {
		s.data = append(s.data, data...)
	} else {
		for i, d := range data {
			if s.size+i < realSize {
				s.data[s.size+i] = d
			} else {
				s.data = append(s.data, data[i:]...)
				break
			}
		}
	}
	s.size += len(data)
}

// Pop 出栈
func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("stack is empty")
	}
	s.size--
	return s.data[s.size], nil
}

// Top 栈顶数据
func (s *Stack) Top() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("stack is empty")
	}
	return s.data[s.size-1], nil
}
