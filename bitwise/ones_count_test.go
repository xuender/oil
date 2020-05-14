package bitwise

import "testing"

func TestBitwise_OnesCount(t *testing.T) {
	var b Bitwise = New()
	if b.OnesCount() != 0 {
		t.Errorf("位计数错误 %d != %d", 0, b.OnesCount())
	}
	b[0] = 3
	if b.OnesCount() != 2 {
		t.Errorf("位计数错误 %d != %d", 2, b.OnesCount())
	}
}
