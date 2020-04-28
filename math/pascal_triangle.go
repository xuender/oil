package math

// PascalTriangle 帕斯卡三角
func PascalTriangle(layer int) []int {
	if layer < 1 {
		return []int{}
	}
	ret := make([]int, layer)
	ret[0] = 1
	if len(ret) == 1 {
		return ret
	}
	old := PascalTriangle(layer - 1)
	for i := 1; i < len(old); i++ {
		ret[i] = old[i] + old[i-1]
	}
	ret[len(old)] = 1
	return ret
}
