# bitfield
### Bit field data structure implemented in Go.

#### Available methods
* Toggle(index) -> toggles the state of a bit at a specific index,
* Get(index) -> retrieve the state of a bit at a specific index.

#### Example usage
```go
package main

import (
  "fmt"
  "github.com/williamsandytoes/bitfield"
)

func main() {
    b := bitfield.New(1000)

    b.Toggle(10)
    fmt.Println(b.Get(10)) // true
    b.Toggle(10)
    fmt.Println(b.Get(10)) // false
}
```
