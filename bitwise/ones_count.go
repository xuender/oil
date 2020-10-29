package bitwise

import "math/bits"

// OnesCount returns the number of one bits.
func (b *Bitwise) OnesCount() int {
	c := 0

	for _, b := range *b {
		c += bits.OnesCount8(b)
	}

	return c
}
