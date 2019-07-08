package dump

import (
	"bytes"
	"encoding/json"
	"fmt"
	"runtime"
	"strings"
)

const (
	COLOR_DEFAULT = "__color__"
	COLOR_YELLOW  = "\033[1;33;40m" + COLOR_DEFAULT + "\033[0m"
	COLOR_GREEN   = "\033[1;32;40m" + COLOR_DEFAULT + "\033[0m"
	COLOR_RED     = "\033[1;31;47m" + COLOR_DEFAULT + "\033[0m"
	COLOR_BLUE    = "\033[1;33;34m" + COLOR_DEFAULT + "\033[0m"
)

var ColorFormat string = COLOR_YELLOW

var Debug bool = true //只有Debug 开启的情况下, 才输出

//与Printf()等价的
func P(format string, v ...interface{}) {
	if Debug {
		fmt.Printf(buildColorFormat(format), newValue(v...)...) //根据设置的颜色, 修改输出
	}
}

//等价P()
func Printf(format string, v ...interface{}) {
	if Debug {
		fmt.Printf(buildColorFormat(format), newValue(v...)...) //根据设置的颜色, 修改输出
	}
}

func newValue(v ...interface{}) []interface{} {
	newV := make([]interface{}, 0, len(v)+1)
	newV = append(newV, callLine(3))
	newV = append(newV, v...)
	return newV
}

func callLine(num int) string {
	pc, file, line, ok := runtime.Caller(num)
	if !ok {
		return ""
	}
	fc := runtime.FuncForPC(pc)
	funcName := ""
	if fc != nil {
		funcName = fc.Name()
		if pos := strings.LastIndex(funcName, "/"); pos > 0 {
			funcName = funcName[pos+1:]
		}
	}
	return fmt.Sprintf("%s:%d:%s()", file[strings.LastIndex(file, "/")+1:], line, funcName)
}

func Println(v ...interface{}) {
	if Debug {
		print(callLine(2), "\t", strings.Replace(ColorFormat, COLOR_DEFAULT, fmt.Sprintln(v...), 1))
	}
}

// 格式化打印JSON
func PrintJsonString(s string) {
	PrintJsonByte([]byte(s))
}

// 格式化打印JSON
func PrintJsonByte(b []byte) {
	var out bytes.Buffer
	err := json.Indent(&out, b, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Println(out.String())
}

func buildColorFormat(format string) string {
	s := "%v\t" + strings.Replace(ColorFormat, COLOR_DEFAULT, format, 1)
	return s
}

//设置颜色, 请传入常量
func SetColor(color string) {
	if strings.Index(color, COLOR_DEFAULT) == -1 {
		panic(fmt.Errorf("color format is error"))
	}
	ColorFormat = color
}

//还原默认颜色
func SetDefaultColor() {
	SetColor(COLOR_YELLOW)
}
