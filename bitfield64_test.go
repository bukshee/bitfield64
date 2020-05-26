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

	if New().Flip(3) != 8 {
		t.Error("shoud be 8")
	}
	if New().Flip(4).Flip(4) != 0 {
		t.Error("should be 0")
	}
}

func Test2(t *testing.T) {
	a := New().Set(1).Set(3)
	if a.Mid(0, 4).OnesCount() != 2 {
		t.Error("should be 2")
	}
	if a.Mid(1, 1).OnesCount() != 1 {
		t.Error("should be 1")
	}
	if a.Mid(3, 5).OnesCount() != 1 {
		t.Error("should be 1")
	}
	if a.Mid(4, 5).OnesCount() != 0 {
		t.Error("should be 0")
	}
	a = New().Set(0).Set(1).Set(4)
	if a.Left(4).OnesCount() != 2 {
		t.Error("should be 2")
	}
	if a.Left(1).OnesCount() != 1 {
		t.Error("should be 1")
	}
	if !New().Set(63).Right(3).Get(2) {
		t.Error("should be true")
	}

	if !a.Rotate(1).Get(5) {
		t.Error("should be true")
	}
	if !a.Rotate(-1).Get(3) {
		t.Error("should be true")
	}
	if !New().Set(5).Rotate(0).Rotate(1).Rotate(-1).Get(5) {
		t.Error("should be true")
	}

	if !New().Set(0).Shift(1).Get(1) {
		t.Error("should be true")
	}
	if !New().Set(1).Shift(-1).Get(0) {
		t.Error("should be true")
	}
	if New().SetAll().Shift(63).OnesCount() != 1 {
		t.Error("should be 1")
	}
	if New().SetAll().Shift(-65).OnesCount() != 0 {
		t.Error("should be 0")
	}
	a = New().Set(0).Set(5).Shift(0).Shift(-1).Shift(1)
	if !a.Get(5) {
		t.Error("should be true")
	}
	if a.Get(0) {
		t.Error("should be false")
	}
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	New().Mid(0, -1)
}
