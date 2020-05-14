package bitwise

// Iteration 迭代
func (b *Bitwise) Iteration(callback func(num int) error) {
	for i := 0; i < len(*b)*_size; i++ {
		if b.In(i) {
			if callback(i) != nil {
				return
			}
		}
	}
}
