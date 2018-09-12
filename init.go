package bitfield

func New(size int) BitField { return BitField{fields: make([]byte, size >> 3)} }
