package integer

// Pow 指数运算
func Pow(x, y int) int {
	r := x
	for i := 1; i < y; i++ {
		r *= x
	}
	return r
}
