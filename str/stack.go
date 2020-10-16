package str

import "errors"

var errStackIsEmpty = errors.New("Stack is Empty")

// Stack 栈
type Stack struct {
	size int
	data []string
}

// NewStack 新建栈
func NewStack(data ...string) *Stack {
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
func (s *Stack) Push(data ...string) {
	if s.IsEmpty() {
		s.data = []string{}
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
func (s *Stack) Pop() (string, error) {
	if s.IsEmpty() {
		return "", errStackIsEmpty
	}
	s.size--

	return s.data[s.size], nil
}

// Top 栈顶数据
func (s *Stack) Top() (string, error) {
	if s.IsEmpty() {
		return "", errStackIsEmpty
	}

	return s.data[s.size-1], nil
}
