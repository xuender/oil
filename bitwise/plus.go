package bitwise

// Plus 加运算
func (b *Bitwise) Plus(num int) {
	bs := make([]byte, _size)

	for i := range bs {
		bs[i] = uint8(num >> (8 * i))
	}

	l := len(*b)

	if l < _size+1 {
		nb := make([]byte, _size-l+1)
		*b = append(*b, nb...)
	}

	for i := range bs {
		b.plus(bs[i], i)
	}
}

func (b *Bitwise) plus(num byte, i int) {
	v := (*b)[i] ^ num
	p := (*b)[i] & num
	(*b)[i] = v + p<<1

	if (v+1) == 255 && p > 0 {
		b.plus(1, i+1)
	}
}
