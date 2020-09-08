package integer

import (
	"fmt"
)

func ExampleBase57() {
	s, _ := Base57.ToString(2020)
	fmt.Println(s)
	s, _ = Base57.ToString(1599526164937)
	fmt.Println(s)

	n, _ := Base57.ToInt("Td")
	fmt.Println(n)
	n64, _ := Base57.ToUint64("mwofPep")
	fmt.Println(n64)

	// Output:
	// Td
	// mwofPep
	// 2020
	// 1599526164937
}

func ExampleBase64() {
	s, _ := Base64.ToString(2020)
	fmt.Println(s)
	s, _ = Base64.ToString(1599526164937)
	fmt.Println(s)

	// Output:
	// kf
	// JnFMrRX
}
