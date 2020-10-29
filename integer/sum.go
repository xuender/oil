package integer

// Sum 求和
func Sum(nums ...int) int {
	ret := 0
	for _, n := range nums {
		ret += n
	}

	return ret
}
