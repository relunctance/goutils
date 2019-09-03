# goutils


## Install 
```shell
go get -u -v github.com/relunctance/goutils
```


### Notice 
- jsondel has been remove to [https://github.com/relunctance/djson](https://github.com/relunctance/djson)

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


