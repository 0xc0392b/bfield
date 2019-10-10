package bfield

import (
  "testing"
)

func TestGet(t *testing.T) {
  b := New(100)

  for i := 0; i < 100; i++ {
    if b.Get(i) != false {
      t.Errorf("Get(%d) = %t; want false", i, b.Get(i))
    }
  }
}

func TestSet(t *testing.T) {
  b := New(100)

  for i := 0; i < 100; i++ {
    b.Set(i, true)
    if b.Get(i) == false {
      t.Errorf("Get(%d) = %t; want true", i, b.Get(i))
    }
  }

  for i := 0; i < 100; i++ {
    b.Set(i, false)
    if b.Get(i) == true {
      t.Errorf("Get(%d) = %t; want false", i, b.Get(i))
    }
  }
}

func TestToggle(t *testing.T) {
  b := New(100)

  for i := 0; i < 100; i++ {
    if b.Get(i) != false {
      t.Errorf("Get(%d) = %t; want false", i, b.Get(i))
    }

    b.Toggle(i)

    if b.Get(i) != true {
      t.Errorf("Get(%d) = %t; want true", i, b.Get(i))
    }

    b.Toggle(i)

    if b.Get(i) != false {
      t.Errorf("Get(%d) = %t; want false", i, b.Get(i))
    }
  }
}
