package bitwise

import (
	"fmt"
	"math/bits"
)

const _size = 8

// Bitwise 位运算
type Bitwise []byte

// Add 增加
func (b *Bitwise) Add(nums ...int) {
	for _, n := range nums {
		size := n/_size + 1
		l := len(*b)
		if l < size {
			nb := make([]byte, size-l)
			*b = append(*b, nb...)
		}
		(*b)[n/_size] |= (1 << (n % _size))
	}
}

// Del 删除
func (b *Bitwise) Del(nums ...int) {
	for _, n := range nums {
		size := n / _size
		if size >= len(*b) {
			continue
		}
		(*b)[size] |= (1 << (n % _size))
		(*b)[size] ^= (1 << (n % _size))
	}
}

// Slice 转换切片
func (b *Bitwise) Slice() []int {
	s := make([]int, b.OnesCount())
	i := 0
	for n := 0; n < len(*b)*_size; n++ {
		if b.In(n) {
			s[i] = n
			i++
		}
	}
	return s
}

// String 转换字符串
func (b *Bitwise) String() string {
	b.Shrink()
	return fmt.Sprintf("%X", *b)
}

// Shrink 收缩
func (b *Bitwise) Shrink() {
	size := len(*b) - 1
	for i := size; i >= 0; i-- {
		if bits.OnesCount8((*b)[i]) > 0 {
			if i == size {
				return
			}
			*b = (*b)[:i+1]
			return
		}
	}
}

// New 新建
func New() Bitwise {
	return make([]byte, _size)
}
