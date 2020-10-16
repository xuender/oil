package math

import "github.com/xuender/oil/integer"

// PascalTriangle 帕斯卡三角非递归
func PascalTriangle(layer int) []int {
	if layer < 0 {
		return []int{}
	}

	ret := make([]int, layer+1)
	ret[0] = 1

	for c := 0; c < layer; c++ {
		ret[c+1] = ret[c] * (layer - c) / (c + 1)
	}

	return ret
}

// PascalSum 帕斯卡三角求和
func PascalSum(layer, num int) int {
	pt := PascalTriangle(layer)

	if num > layer {
		return integer.Sum(pt...)
	}

	return integer.Sum(pt[num:]...)
}
