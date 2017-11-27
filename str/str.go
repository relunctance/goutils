package str

import (
	"fmt"
	"math"
	"strconv"
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
