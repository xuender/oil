package bitwise

import (
	"testing"
)

func TestBitwise_In(t *testing.T) {
	var b Bitwise = New()
	b.Add(0, 1, 2, 3)
	if b.OnesCount() != 4 {
		t.Errorf("位计数错误 %d != %d", 4, b.OnesCount())
	}
	if !b.In(3) {
		t.Errorf("In 错误 %d", 3)
	}
	if b.In(4) {
		t.Errorf("In 错误 %d", 4)
	}
	if b.InAny(4, 5, 6) {
		t.Error("InAny 错误 ")
	}
	if !b.InAny(1, 2, 6) {
		t.Error("InAny 错误 ")
	}
	if !b.InAll(1, 2) {
		t.Error("InAll 错误 ")
	}
	if b.InAll(1, 2, 6) {
		t.Error("InAll 错误 ")
	}
}
