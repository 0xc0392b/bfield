package bfield

// Toggle toggles the bit's state at index i in the bit field's
// byte array
func (f *Field) Toggle(i int) {
  f.vec[i >> 3] ^= 1 << uint(i & 7)
}

// Get returns true if the bit at index i in the bit field's
// byte array is 1, else it returns false
func (f *Field) Get(i int) bool {
  return f.vec[i >> 3] & (1 << uint((i & 7))) != 0
}

// Set sets the state of the bit at index i in the bit field's
// byte array to 1 if state is true else 0
func (f *Field) Set(i int, state bool) {
  if state == true {
    f.vec[i >> 3] |= 1 << uint(i & 7)
  } else {
    f.vec[i >> 3] &= (255 ^ 1 << uint(i & 7))
  }
}

// Len returns the number of bits in the underlying byte array
// - which will be (length of the array * 8) thus Len will never
// return the original size unless (size % 8 = 0)
func (f *Field) Len() int {
  return len(f.vec) * 8
}
