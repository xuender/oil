package array

import "fmt"

func ExampleNewProduct() {
	p, _ := NewProduct(2, 3)
	for p.Next() {
		fmt.Println(p.Value())
	}

	// Output:
	// [0 0 0]
	// [0 0 1]
	// [0 1 0]
	// [0 1 1]
	// [1 0 0]
	// [1 0 1]
	// [1 1 0]
	// [1 1 1]
}
