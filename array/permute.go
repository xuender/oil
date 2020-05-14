package array

import "sort"

// Permute Lexicographic 字典序法,排列
func Permute(n int, callback func([]int) error) {
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i
	}

	for {
		if callback(nums) != nil {
			return
		}

		pos := -1
		for i := n - 1; i > 0; i-- {
			if nums[i] > nums[i-1] {
				pos = i
				break
			}
		}
		if pos == -1 {
			break
		}

		sort.Ints(nums[pos:])
		// swap with the smallest one which is bigger than nums[pos]
		for i := pos; i < n; i++ {
			if nums[i] > nums[pos-1] {
				nums[i], nums[pos-1] = nums[pos-1], nums[i]
				break
			}
		}
	}
}
