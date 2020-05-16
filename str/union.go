package str

// Union 字符串转换[]byte，byte唯一
func Union(str string) []byte {
	bytes := []byte(str)
	m := map[byte]bool{}
	for _, b := range bytes {
		m[b] = true
	}

	bs := make([]byte, len(m))
	i := 0
	for _, b := range bytes {
		if ok := m[b]; ok {
			bs[i] = b
			i++
			m[b] = false
		}
	}

	return bs
}
