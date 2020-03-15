package integer

// Div 除法，四舍五入
func Div(dividend, divisor int) int {
	res := dividend * 10 / divisor
	if res%10 > 4 {
		return res/10 + 1
	}
	return res / 10
}
