package integer

import (
	"fmt"
	"unsafe"
)

// Encoding base
type Encoding struct {
	encode []byte
}

// NewEncoding returns a padded Encoding defined by the given alphabet
func NewEncoding(encoder string) *Encoding {
	if len(encoder) < 2 {
		panic("encoding alphabet too short")
	}
	for i := 0; i < len(encoder); i++ {
		if encoder[i] == '\n' || encoder[i] == '\r' {
			panic("encoding alphabet contains newline character")
		}
	}
	encode := make([]byte, len(encoder))
	copy(encode, encoder)
	e := &Encoding{encode: encode}
	return e
}

// ToString returns number to string
func (enc *Encoding) ToString(num interface{}) (string, error) {
	var n uint64
	switch num.(type) {
	case byte:
		n = uint64(num.(byte))
	case uintptr:
		n = uint64(num.(uintptr))
	case int:
		n = uint64(num.(int))
	case int8:
		n = uint64(num.(int8))
	case int16:
		n = uint64(num.(int16))
	case int32:
		n = uint64(num.(int32))
	case uint:
		n = uint64(num.(uint))
	case uint16:
		n = uint64(num.(uint16))
	case uint32:
		n = uint64(num.(uint32))
	case uint64:
		n = num.(uint64)
	default:
		return "", fmt.Errorf("type not number")
	}
	return enc.uint64ToString(n), nil
}

func (enc *Encoding) uint64ToString(num uint64) string {
	ret := []byte{}
	l := uint64(len(enc.encode))
	for num > 0 {
		ret = append(ret, enc.encode[num%l])
		num = num / l
	}
	return string(ret)
}
func (enc *Encoding) index(b byte) (uint64, error) {
	for i, char := range enc.encode {
		if char == b {
			return uint64(i), nil
		}
	}
	return 0, fmt.Errorf("Element '%v' is not part of the alphabet", b)
}

// ToUint64 returns string to uint64
func (enc *Encoding) ToUint64(str string) (uint64, error) {
	var n uint64
	l := uint64(len(enc.encode))

	for i := len(str) - 1; i >= 0; i-- {
		n = n * l
		index, err := enc.index(str[i])
		if err != nil {
			return 0, err
		}
		n += index
	}
	return n, nil
}

// ToInt64 returns string to int64
func (enc *Encoding) ToInt64(str string) (int64, error) {
	n, err := enc.ToUint64(str)
	if err != nil {
		return 0, err
	}
	return int64(n), nil
}

// ToInt returns string to int
func (enc *Encoding) ToInt(str string) (int, error) {
	n, err := enc.ToUint64(str)
	if err != nil {
		return 0, err
	}
	intNum := *(*int)(unsafe.Pointer(&n))
	return intNum, nil
}
