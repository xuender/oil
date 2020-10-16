package random

import (
	"math/rand"
)

// Queue 随机队列
type Queue []int

// NewQueue 新建随机队列
func NewQueue(max, size int) Queue {
	if size >= max {
		ret := make(Queue, max)
		for i := 0; i < max; i++ {
			ret[i] = i
		}

		return ret
	}

	ints := make(Queue, size)
	m := map[int]bool{}

	for {
		m[rand.Intn(max)] = true
		if len(m) >= size {
			break
		}
	}

	i := 0

	for num := range m {
		ints[i] = num
		i++
	}
	// 单一职责 不应排序
	// sort.Ints(ints)
	return ints
}
