package main

import (
	"fmt"

	"github.com/xuender/oil/permute"
)

func main() {
	permute.Lexicographic(3, func(bs []int) error {
		fmt.Println(bs)
		return nil
	})
}
