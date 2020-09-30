package array

import "fmt"

func ExampleExchange() {
	e := NewExchange([]int{2, 3, 2})
	for e.Next() {
		fmt.Println(e.Value())
	}

	// Output:
	// [0 0 0]
	// [0 0 1]
	// [0 1 0]
	// [0 1 1]
	// [0 2 0]
	// [0 2 1]
	// [1 0 0]
	// [1 0 1]
	// [1 1 0]
	// [1 1 1]
	// [1 2 0]
	// [1 2 1]
}
