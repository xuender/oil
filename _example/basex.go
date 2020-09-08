package main

import (
	"fmt"

	"github.com/xuender/oil/basex"
)

func main() {
	for i := 0; i < 1000; i++ {
		a, _ := basex.Base57.ToString(i)
		b, _ := basex.Base57.ToInt(a)

		c, _ := basex.Base64.ToString(i)
		d, _ := basex.Base64.ToInt(c)
		fmt.Println(a, "=", b, ";", c, "=", d)
	}
}
