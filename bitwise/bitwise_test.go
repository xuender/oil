package bitwise

import (
	"testing"
)

func TestBitwise_Add(t *testing.T) {
	b := New()
	b.Add(1024)
	if len(b) < 100 {
		t.Errorf("增加后长度不对:%d", len(b))
	}
	if b.OnesCount() != 1 {
		t.Errorf("增加后数量不对:%d", b.OnesCount())
	}
	b.Add(1)
	b.Add(2)
	b.Add(2)
	if b.OnesCount() != 3 {
		t.Errorf("增加后数量不对:%d", b.OnesCount())
	}
}

func TestBitwise_Del(t *testing.T) {
	b := New()
	b.Add(1, 2, 3)
	if b.OnesCount() != 3 {
		t.Errorf("增加后数量不对:%d", b.OnesCount())
	}
	b.Del(2, 3, 4)
	if b.OnesCount() != 1 {
		t.Errorf("删除后数量不对:%d", b.OnesCount())
	}
}

func TestBitwise_Slice(t *testing.T) {
	b := New()
	b.Add(1, 2, 3)
	if b.Slice()[1] != 2 {
		t.Error("切片错误d")
	}
}

func TestBitwise_String(t *testing.T) {
	b := New()
	b.Add(12)
	b.Shrink()
	if !b.In(12) {
		t.Error("Reset错误")
	}
	if b.String() != "0010" {
		t.Errorf("String 错误 %s", b.String())
	}
	b.Add(32)
	b.Del(32)
	if b.String() != "0010" {
		t.Errorf("String 错误 %s", b.String())
	}
}
