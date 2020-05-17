# bitfield64
Bitfield for up to 64bits in length

## Description
Package bitfield64 is a simple, quick stack-based bit-field manipulator
package of 64 bits (or less) in length. If you need more bits you either
need to create an array of bitfield64-s (and stay on the stack) or need to
switch to the heap-based bitfield package. Methods are stateless and free
from side-effects. It was designed to be chainable so error reporting is
omitted: you need to make sure position is in the range of [0, 63]

For usage see test file.

## Badges

![License](https://img.shields.io/github/license/bukshee/bitfield64)
![Version](https://img.shields.io/github/manifest-json/v/bukshee/bitfield64)
![Downloads](https://img.shields.io/github/downloads/bukshee/bitfield64/total)
