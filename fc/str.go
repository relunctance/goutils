package fc

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"math"
	"net/url"
	"strconv"
	"strings"
)

func Chr(c rune) string {
	return string(c)
}

func Ord(c string) []rune {
	return []rune(c)
}

func Substr(str string, begin, length int) string {
	lth := len(str)
	if begin < 0 {
		begin = 0
	}
	if begin >= lth {
		begin = lth
	}
	end := begin + length
	if end > lth {
		end = lth
	}

	return string(str[begin:end])
}

func StrToInt32(str string) (int32, error) {
	num, err := strconv.ParseInt(str, 10, 0)
	if err != nil {
		return 0, fmt.Errorf("\"%s\" is not integer", str)
	}
	if num > math.MaxInt32 || num < math.MinInt32 {
		return 0, fmt.Errorf("%d is not 32-bit integer", num)
	}
	return int32(num), nil
}

/* eg:
u := ParseStr("?username=zhangsan&passwd=123456")
u.Get("username");
*/
func ParseStr(str string) url.Values {
	if strings.Index(str, "?") == -1 {
		panic(fmt.Errorf("must exists '?' like: '?username=zhangsan&passwd=123456'"))
	}
	u, err := url.Parse(str)
	if err != nil {
		panic(err)
	}
	return u.Query()
}

//首字母大写
func Ucfirst(str string) string {
	if len(str) == 0 {
		return str
	}
	if len(str) == 1 {
		return strings.ToUpper(str)
	}
	return string(bytes.ToUpper([]byte{str[0]})) + str[1:]
}

// md5 计算
func Md5(str string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}

func Md5Bytes(v []byte) string {
	return fmt.Sprintf("%x", md5.Sum(v))
}

func ByteFormat(i float64) string {
	var a = []string{"B", "KB", "MB", "GB", "TB", "PB", "EB", "ZB", "YB", "UnKnown"}
	var pos int = 0
	var j float64 = float64(i)
	for {
		if i >= 1024 {
			i = i / 1024
			j = j / 1024
			pos++
		} else {
			break
		}
	}
	if pos >= len(a) { // fixed out index bug
		pos = len(a) - 1
	}
	return fmt.Sprintf("%.3f %s", j, a[pos])
}
