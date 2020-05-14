package bitwise

import (
	"fmt"
)

func ExampleBitwise_Iteration() {
	b := New()
	b.Add(1, 3, 4, 99)
	b.Del(3)
	b.Iteration(func(i int) error {
		fmt.Println(i)
		return nil
	})

	// Output:
	// 1
	// 4
	// 99
}
