package debug

import (
	"fmt"
	"runtime"
	"strings"
)

//返回tracedebug 函数调用栈路径信息
//文件名+行号+函数名
func DebugTrace() []string {
	i := 0
	ret := make([]string, 0, 10)
	funcName := ""
	for {
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		fc := runtime.FuncForPC(pc)
		if fc != nil {
			funcName = fc.Name()
			if pos := strings.LastIndex(funcName, "/"); pos > 0 {
				funcName = funcName[pos+1:]
			}
		}
		ret = append(ret, fmt.Sprintf("%s:%d:%s()\n", file, line, funcName))
		i++
	}
	return ret
}
