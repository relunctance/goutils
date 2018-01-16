/*
一维数组:  获取边界
*/
package slicelimit

import (
	"errors"
	"fmt"
)

var EOF error = errors.New("END")

//根据offset和pagesize获取边界
func GetBoundary(offset, pagesize, length int64) (err error, start int64, end int64) {

	if offset > length {
		err = fmt.Errorf("offset:[%d] has exceed array boundary:[%d] ", offset, length)
		return
	}

	width := offset + pagesize
	if width > length { // 当需要取的范围超过 数组最大长度时 , 直接返回从数组最后一位
		start = offset
		end = length
		if start == end {
			err = EOF //结束标识符
		}
		return
	}
	start = offset
	end = width
	return
}

//检测是否结束
func CheckIsEnd(err error) bool {
	if err == EOF {
		return true
	}
	if err != nil {
		panic(err)
	}
	return false
}
