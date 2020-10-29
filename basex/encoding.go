package basex

import (
	"errors"
	"unsafe"

	"github.com/xuender/oil/str"
)

var (
	errTypeNumber = errors.New("Type not number")
	errElement    = errors.New("Element not part of the alphabet")
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

	e := &Encoding{encode: str.Union(encoder)}

	return e
}

// ToString returns number to string
func (enc *Encoding) ToString(num interface{}) (string, error) {
	var n uint64
	switch num := num.(type) {
	case byte:
		n = uint64(num)
	case uintptr:
		n = uint64(num)
	case int:
		n = uint64(num)
	case int8:
		n = uint64(num)
	case int16:
		n = uint64(num)
	case int32:
		n = uint64(num)
	case int64:
		n = uint64(num)
	case uint:
		n = uint64(num)
	case uint16:
		n = uint64(num)
	case uint32:
		n = uint64(num)
	case uint64:
		n = num
	default:
		return "", errTypeNumber
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

	if len(ret) == 0 {
		ret = append(ret, enc.encode[0])
	}

	return string(ret)
}
func (enc *Encoding) index(b byte) (uint64, error) {
	for i, char := range enc.encode {
		if char == b {
			return uint64(i), nil
		}
	}

	return 0, errElement
}

// ToUint64 returns string to uint64
func (enc *Encoding) ToUint64(str string) (uint64, error) {
	var (
		n uint64
		l = uint64(len(enc.encode))
	)

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
