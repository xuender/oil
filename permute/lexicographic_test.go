package permute

import "fmt"

func ExampleLexicographic() {
	Lexicographic(3, func(array []int) error {
		fmt.Println(array)
		return nil
	})
	// Output:
	// [0 1 2]
	// [0 2 1]
	// [1 0 2]
	// [1 2 0]
	// [2 0 1]
	// [2 1 0]
}
