package array

import (
	"github.com/xuender/oil/bitwise"
	"github.com/xuender/oil/math"
)

// Combine 组合
func Combine(num int, callback func([]int) error) {
	b := bitwise.New()
	for i := 0; i < math.Pow(2, num)-1; i++ {
		b.Plus(1)
		if callback(b.Slice()) != nil {
			return
		}
	}
}
