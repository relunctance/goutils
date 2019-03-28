package fc

import (
	"encoding/hex"
	"fmt"
	"os"
)

func Hex2bin(raw string) string {
	result, _ := hex.DecodeString(raw)
	return string(result)
}

func Bin2hex(raw string) string {
	return hex.EncodeToString([]byte(raw))
}

func HexDump(raw string) string {
	v := Hex2bin(s)
	return hex.Dump([]byte(v))
}

/*
Returns a string with backslashes added before characters that need to be escaped. These characters are:

single quote (')
double quote (")
backslash (\)
*/
func Addslashes(str string) string {
	ret := make([]rune, 0, len(str))
	for _, c := range str {
		switch c {
		case
			'\\',
			'"',
			'\'':
			ret = append(ret, '\\')
		}
		ret = append(ret, c)
	}
	return string(ret)
}

// Un-quotes a quoted string.
func Stripslashes(str string) string {
	l := len(str)
	ret := make([]rune, 0, l)
	for i := 0; i < l; i++ {
		if str[i] == '\\' {
			i++
		}
		ret = append(ret, rune(str[i]))
	}
	return string(ret)
}

// 使用当前时间纳秒+hostname+pid的唯一ID
// prefix用于解决结果碰撞问题
// length =  26
func Uniqid(prefix string) string {
	f := Random()
	h, _ := os.Hostname()
	h = Md5(fmt.Sprintf("%s-%d-%f", h, os.Getpid(), f))
	fs := fmt.Sprintf("%f", f)
	return prefix + h[0:20] + fs[2:]
}
