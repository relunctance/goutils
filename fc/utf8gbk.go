package fc

import (
	"unicode/utf8"

	"github.com/axgle/mahonia"
)

// IsCanUtf8ToGbk  可以针对一些unicode 转换后的内容进一步判断,是否还需要再做一次转换
func IsCanUtf8ToGbk(s string) bool {
	if len(s) == 0 {
		return false
	}
	return IsUtf8(Utf8ToGbk(s))
}

// IsUtf8 判断是否是utf8
func IsUtf8(s string) bool {
	return utf8.ValidString(s)
}

func GbkToUtf8(v string) string {
	dec := mahonia.NewDecoder("gbk") //decode
	ret := dec.ConvertString(v)
	return ret

}

func Utf8ToGbk(v string) string {
	enc := mahonia.NewEncoder("gbk") //encode
	ret := enc.ConvertString(v)
	return ret
}
