package integer

// Max 取最大值
func Max(nums ...int) int {
	max := MinInt
	for _, n := range nums {
		if max < n {
			max = n
		}
	}
	return max
}
