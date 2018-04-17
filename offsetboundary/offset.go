/*
一维数组:  获取边界
*/
package offsetboundary

import (
	"errors"
	"fmt"
)

var EOF error = errors.New("END")

//根据offset和pagesize获取边界
//函数功能请见测试用例
func GetBoundary(offset, pagesize, length int64) (err error, start int64, end int64) {

	if offset > length {
		err = fmt.Errorf("offset:[%d] has exceed array boundary:[%d] ", offset, length)
		return
	}

	if offset == length { // 边界的情况
		err = EOF
		return
	}

	start = offset

	if pagesize >= length {
		end = length
		err = EOF
		return

	}

	width := offset + pagesize
	if width >= length { // 当需要取的范围超过 数组最大长度时 , 直接返回从数组最后一位
		end = length
		err = EOF //结束标识符
		return
	}
	end = width
	return
}

//检测是否结束
func checkIsEnd(err error) (bool, error) {
	if err == EOF {
		return true, nil
	}
	if err != nil {
		return true, err
	}
	return false, nil
}

//获取HasNext
func GetHasNext(err error) (bool, error) {
	hasNext, err2 := checkIsEnd(err)
	hasNext = !hasNext
	return hasNext, err2
}
