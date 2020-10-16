package integer

// Ints 字符串转换正整数数组
func Ints(str string) []int {
	res := make([]int, len(str))

	for i, s := range str {
		res[i] = int(s - 48)
	}

	return res
}
