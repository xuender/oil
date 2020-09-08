package integer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewEncoding(t *testing.T) {
	e := NewEncoding("123")
	assert.Equal(t, len(e.encode), 3, "len")
	assert.Equal(t, e.encode[1], uint8('2'), "2")
}

func TestEncoding_ToString(t *testing.T) {
	e := NewEncoding("123")
	s, _ := e.ToString(5)
	assert.Equal(t, s, "32", "len")
}

func TestEncoding_ToUint64(t *testing.T) {
	e := NewEncoding("123")
	s, _ := e.ToUint64("32")
	assert.Equal(t, s, uint64(5), "len")
}
