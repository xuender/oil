package integer

import "fmt"

func ExampleNewStack() {
	s := NewStack(1)
	s.Push(2)
	fmt.Println(s.Pop())
	s.Push(3, 4, 5)
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Size())
	fmt.Println(s.Top())
	s.Push(9)
	fmt.Println(s.Top())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())

	// Output:
	// 2 <nil>
	// 5 <nil>
	// 4 <nil>
	// 3 <nil>
	// 1 <nil>
	// 0
	// 0 stack is empty
	// 9 <nil>
	// 9 <nil>
	// 0 stack is empty
}
