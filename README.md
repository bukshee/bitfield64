![Coverage](https://img.shields.io/badge/coverage-100%25-green)

# bitfield64
Bitfield for up to 64bits in length

## Description
Package bitfield64 is a simple, quick stack-based bit-field manipulator
package of 64 bits (or less) in length. If you need more
bits you either need to create an array of bitfield64-s (and stay on the stack)
or need to switch to the heap-based bitfield package.

Methods are stateless and free from side-effects.

It was designed to be chainable. Range for position must be [0, 63]. position
outside this range will get the modulo treatment, so 64 will point to the 0th
element, -1 will address the last element (i.e. 63rd), -2 the one before
(i.e. 62nd), etc.

For usage see test file.
