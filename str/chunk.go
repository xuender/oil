package str

import (
	"github.com/xuender/oil/integer"
)

// Chunk 分组
func Chunk(str string, size int) []string {
	length := len(str)
	if length == 0 || size < 1 {
		return []string{}
	}

	ret := make([]string, integer.Ceil(length, size))
	index := 0
	i := 0

	for index < length {
		end := index + size
		if end > length {
			end = length
		}

		ret[i] = str[index:end]
		index = end
		i++
	}

	return ret
}
