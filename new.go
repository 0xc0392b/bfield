package bfield

import (
  "math"
)

// New returns a pointer to an instance of Field - which contains
// a byte array (vector) of length (⌈size / 8⌉)
// Ceil accounts for size not dividing equally into 8 which would
// result in Len() < size since the decimal places are ignored
func New(size int) *Field {
  return &Field{vec: make([]byte, int32(math.Ceil(float64(size) / 8)))}
}
