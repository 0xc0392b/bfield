package bfield

import (
  "testing"
)

func TestNew(t *testing.T) {
  var b *Field

  b = New(8)
  if b.Len() != 8 {
    t.Errorf("New(8).Len() = %d; want 8", b.Len())
  }

  b = New(1010)
  if b.Len() != 1016 {
    t.Errorf("New(1000).Len() = %d; want 1016", b.Len())
  }
}
