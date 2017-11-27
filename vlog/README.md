## 用法示例:

```go
package main                                                                                                         

import (
        "fmt"                                                                                                            
        "os"

        "github.com/relunctance/goutils/vlog"                                                                            
        "github.com/relunctance/goutils/vlog/base"                                                                       
       )

func main() {
        //用法1:                                                                                                         
        vl := vlog.Logger(base.TYPE_LOGRUS, base.LEVEL_INFO, base.FORMAT_TEXT, os.Stdout, nil)                           
        m := make(map[string]string)
        m["demo1"] = "2"                                                                                                 
        m["demo2"] = "3"                                                                                                 
        vl.Infoln("aaaaaaaaaa", "xxx", string([]byte("hello\t")), m)                                                       
        fmt.Println(vl)
}   
```
