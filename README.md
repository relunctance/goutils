# goutils


## Install 
```shell
go get -u -v github.com/relunctance/goutils
```

### Useage 
```go
package main

import (
    "fmt"

    "github.com/relunctance/goutils/fc"
)

func main() {
    fmt.Println(fc.DebugTrace())
}
```

### Modules

* fc  - Commonly used encoding functions
* dump - Print/Println with color output stdout
* cmd - Execute command with golang 
* offsetboundary - Slice data clipping 


