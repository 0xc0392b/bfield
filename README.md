# bfield
### Bit field data structure for Go

#### Methods
* Toggle(index) -> toggles the state of a bit at a specific index
* Get(index) -> retrieve the state of a bit at a specific index
* Set(index, state) -> sets the state of a bit at a specific index

#### Installation
```text
go get github.com/williamsandytoes/bfield
```

#### Example usage
```go
package main

import (
  "fmt"
  "github.com/williamsandytoes/bfield"
)

func main() {
    b := bfield.New(1000)

    b.Toggle(10)
    fmt.Println(b.Get(10)) // true
    b.Toggle(10)
    fmt.Println(b.Get(10)) // false
}
```
