package bitwise

import (
	"fmt"
)

func ExampleBitwise_Plus() {
	b := New()
	b.Plus(8*8*8*8*8*2 - 1)
	b.Shrink()
	fmt.Println(b)

	b.Plus(1)
	b.Shrink()
	fmt.Println(b)

	// Output:
	// [255 255]
	// [0 0 1]
}
