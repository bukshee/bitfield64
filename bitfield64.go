/*
Package bitfield64 is a simple, quick stack-based bit-field manipulator
package of 64 bits (or less) in length. If you need more
bits you either need to create an array of bitfield64-s (and stay on the stack)
or need to switch to the heap-based bitfield package.

Methods are stateless and free from side-effects.

It was designed to be chainable. Range for position must be [0, 63]. position
outside this range will get the modulo treatment, so 64 will point to the 0th
element, -1 will address the last element (i.e. 63rd), -2 the one before
(i.e. 62nd), etc.
*/
package bitfield64

import (
	"math"
	"math/bits"
)

// BitField64 type utilizing the power of 64bit CPUs. It lives on the stack
type BitField64 uint64

// New returns a zeroed (all false) bit-field that can store size elements
func New() BitField64 {
	return BitField64(0)
}

// pos => bit mask: 1 at bit position pos, 0 everywhere else
func posToBitMask(pos int) BitField64 {
	for pos < 0 {
		pos += 64
	}
	return 1 << uint64(pos%64)
}

// Set sets the bit at position pos
func (bf64 BitField64) Set(pos int) BitField64 {
	return bf64 | posToBitMask(pos)
}

// Get returns true if bit at position pos is set, false otherwise
func (bf64 BitField64) Get(pos int) bool {
	return bf64&posToBitMask(pos) > 0
}

// Clear clears the bit at position pos
func (bf64 BitField64) Clear(pos int) BitField64 {
	return bf64 & ^posToBitMask(pos)
}

// Flip inverts the bit at position pos
func (bf64 BitField64) Flip(pos int) BitField64 {
	if bf64.Get(pos) {
		return bf64.Clear(pos)
	}
	return bf64.Set(pos)
}

// SetAll returns a bitfield where all 64 bits are set
func (bf64 BitField64) SetAll() BitField64 {
	return math.MaxUint64
}

// ClearAll returns a bitfield where all 64 bits are set
func (bf64 BitField64) ClearAll() BitField64 {
	return BitField64(0)
}

// And returns the binary AND of the two bitfields
func (bf64 BitField64) And(bfo BitField64) BitField64 {
	return bf64 & bfo
}

// Or returns the binary OR of the two bitfields
func (bf64 BitField64) Or(bfo BitField64) BitField64 {
	return bf64 | bfo
}

// Not returns the bitfield with each bit inverted: 0 becomes 1, 1 becomes 0
func (bf64 BitField64) Not() BitField64 {
	return ^bf64
}

// Xor returns the binary XOR of the two bitfields
func (bf64 BitField64) Xor(bfo BitField64) BitField64 {
	return bf64 ^ bfo
}

// OnesCount returns the number of bits set
func (bf64 BitField64) OnesCount() int {
	return bits.OnesCount64(uint64(bf64))
}
