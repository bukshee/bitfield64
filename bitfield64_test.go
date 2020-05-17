package bitfield64

import (
	"testing"
)

func TestPos(t *testing.T) {
	got := []int{0, -1, -64, 3, 65}
	wants := []uint64{1 << 0, 1 << 63, 1 << 0, 1 << 3, 1 << 1}
	for i, p := range got {
		if posToBitMask(p) != BitField64(wants[i]) {
			t.Errorf("got %d, wants %d", p, wants[i])
		}
	}
}

func Test1(t *testing.T) {
	bf1 := New().SetAll().Clear(3)
	if bf1.OnesCount() != 63 {
		t.Error("should be 63")
	}
	if bf1.Not().OnesCount() != 1 {
		t.Error("should be 1")
	}
	if bf1.Get(3) {
		t.Error("should be false")
	}
	if !bf1.Get(17) {
		t.Error("should be true")
	}
	bf2 := New().Set(3).Set(1).Clear(5).Or(New().Set(4))

	if bf2.OnesCount() != 3 {
		t.Error("should be 3")
	}
	if bf1.Xor(bf1).OnesCount() != 0 {
		t.Error("should be zero")
	}
	if bf2.And(bf1).OnesCount() != 2 {
		t.Error("should be 2")
	}

	if New().SetAll().ClearAll().OnesCount() != 0 {
		t.Error("should be zero")
	}
	if New().Set(-63-64) != 2 {
		t.Error("should be 2")
	}
}
