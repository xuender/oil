package integer

// Ceil 除法向上进位
func Ceil(dividend, divisor int) int {
	if dividend%divisor > 0 {
		return dividend/divisor + 1
	}

	return dividend / divisor
}
