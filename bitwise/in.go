package bitwise

// In 包含
func (b *Bitwise) In(num int) bool {
	size := num / _size
	if size >= len(*b) {
		return false
	}
	return (*b)[size]&(1<<(num%_size)) > 0
}

// InAny 包含任意个
func (b *Bitwise) InAny(nums ...int) bool {
	for _, num := range nums {
		if b.In(num) {
			return true
		}
	}
	return false
}

// InAll 包含全部
func (b *Bitwise) InAll(nums ...int) bool {
	for _, num := range nums {
		if !b.In(num) {
			return false
		}
	}
	return true
}
