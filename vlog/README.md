## 功能:
支持同时向console和文件记录
支持日志记录具体文件位置和行号
支持TEXT和JSON两种格式输出
## 用法示例:

```go
package main

import (
        "io"
        "os"

        "github.com/relunctance/goutils/vlog"
        "github.com/relunctance/goutils/vlog/base"
       )

func main() {

    f, _ := os.OpenFile("./logtmp.log", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0766)
        multiW := io.MultiWriter(os.Stdout, f) //终端和文件同时输出内容
        defer f.Close()
        option := []base.Option{base.OptWithLocation{true}}
    //用法1:
vl := vlog.Logger(base.TYPE_LOGRUS, base.LEVEL_INFO, base.FORMAT_TEXT, multiW, option)
        m := make(map[string]string)
        m["demo1"] = "2"
        m["demo2"] = "3"
        vl.Infoln("aaaaaaaaaa", "xxx", string([]byte("hello")), m)
}

```
### 执行结果示例:
```
   INFO[2017-11-28 10:44:34.57] aaaaaaaaaa xxx hello        map[demo2:3 demo1:2]    location="map[func_path:main.main file_name:demo.go line:18]"
```
