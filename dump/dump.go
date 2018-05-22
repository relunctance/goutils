package dump

import "fmt"
import "strings"

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
		fmt.Printf(buildColorFormat(format), v...) //根据设置的颜色, 修改输出
	}
}

//等价P()
func Printf(format string, v ...interface{}) {
	if Debug {
		fmt.Printf(buildColorFormat(format), v...) //根据设置的颜色, 修改输出
	}
}

func Println(v ...interface{}) {
	if Debug {
		print(strings.Replace(ColorFormat, COLOR_DEFAULT, fmt.Sprintln(v...), 1))
	}
}

func buildColorFormat(format string) string {
	return strings.Replace(ColorFormat, COLOR_DEFAULT, format, 1)
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
