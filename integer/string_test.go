package integer

import "fmt"

func ExampleString() {
	fmt.Println(String(1, 2, 3, -4))

	// Output:
	// 123-4
}
