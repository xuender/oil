package str

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTime(t *testing.T) {
	assert.Equal(t, "1纳秒", Time(int64(1)))
	assert.Equal(t, "10纳秒", Time(int64(10)))
	assert.Equal(t, "20纳秒", Time(int64(20)))
	assert.Equal(t, "200纳秒", Time(int64(200)))
	assert.Equal(t, "1微妙", Time(int64(1000)))
	assert.Equal(t, "10微妙", Time(int64(10000)))
	assert.Equal(t, "20微妙", Time(int64(20000)))
	assert.Equal(t, "1.23微妙", Time(int64(1234)))
	assert.Equal(t, "10微妙", Time(int64(10000)))
	assert.Equal(t, "200微妙", Time(int64(200000)))
	assert.Equal(t, "1毫秒", Time(int64(1000000)))
	assert.Equal(t, "1.23毫秒", Time(int64(1234567)))
	assert.Equal(t, "10毫秒", Time(int64(10000000)))
	assert.Equal(t, "100毫秒", Time(int64(100000000)))
	assert.Equal(t, "1秒", Time(int64(1000000000)))
	assert.Equal(t, "10秒", Time(int64(10000000000)))
	assert.Equal(t, "1.23秒", Time(int64(1234567890)))
	assert.Equal(t, "1分钟", Time(int64(60000000000)))
	assert.Equal(t, "1分钟1.2秒", Time(int64(61200000000)))
	assert.Equal(t, "1小时10分钟1.2秒", Time(int64(4201200000000)))
	assert.Equal(t, "1天3小时16分钟41.2秒", Time(int64(98201200000000)))
}
