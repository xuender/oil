package integer

// Exp 指数运算
func Exp(num, n int) int {
	res := 1
	for i := n; i > 0; i >>= 1 {
		if i&1 != 0 {
			res *= num
		}
		num *= num
	}
	return res
}
