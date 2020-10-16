package integer

// Min 取最小值
func Min(nums ...int) int {
	min := MaxInt
	for _, n := range nums {
		if min > n {
			min = n
		}
	}

	return min
}
