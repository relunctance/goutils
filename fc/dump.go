package fc

import "fmt"

// 类似PHP  var_dump
func Dump(vals ...interface{}) {
	for _, v := range vals {
		fmt.Println(JsonDump(v))
	}
}
