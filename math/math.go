package math

import (
	"fmt"
	"strconv"
)

//返回较小的值
func MinInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}

//返回较小的值
func MinInt64(x, y int64) int64 {
	if x < y {
		return x
	}
	return y
}

func Round(value float64, precision int) (f float64) {
	if precision < 0 {
		panic(fmt.Errorf("round precision must gt 0"))
	}
	format := fmt.Sprintf("%df", precision)
	s := fmt.Sprintf("%0."+format, value)
	f, _ = strconv.ParseFloat(s, 64)
	return
}
