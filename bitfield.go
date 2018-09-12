package bitfield

func (b *BitField) Toggle(i int) { b.vec[i >> 3] ^= 1 << (i & 7) }

func (b *BitField) Get(i int) bool { return  b.vec[i >> 3] & (1 << (i & 7)) == 0 }
