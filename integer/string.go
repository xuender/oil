package integer

import (
	"strconv"
	"strings"
)

// String 整数数组转换字符串
func String(ints ...int) string {
	l := len(ints)

	if l == 0 {
		return ""
	}

	if l == 1 {
		return strconv.Itoa(ints[0])
	}

	s := make([]string, len(ints))

	for i, n := range ints {
		s[i] = strconv.Itoa(n)
	}

	return strings.Join(s, "")
}
