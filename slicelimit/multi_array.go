package slicelimit

import (
	"fmt"
)

//多维数组使用offset和pagesize
func GetDyadicArrayByOffset(offset, pagesize, totalNum int64, multiDyadicArr [][]int) (error, [][]int) {
	var err error
	var resultArr [][]int = make([][]int, len(multiDyadicArr))
	if offset > totalNum {
		err = fmt.Errorf("offset:[%d] has exceed DyadicArray TotalLength:[%d]  boundary", offset, totalNum)
	}
	errstr, start, _ := GetBoundary(offset, pagesize, totalNum)
	if errstr != nil {
		err = errstr
		return err, resultArr
	}

	var psize int64 = pagesize
	k := int64(0)
	for j, values := range multiDyadicArr {
		//fmt.Println("values:", values)
		vl := int64(len(values)) //3
		//resultArr[j] = make([]int, vl)
		resultArr[j] = nil
		if start > vl {
			start = start - vl
			continue
		}

		_err, _start, _end := GetBoundary(start, psize, vl)
		//fmt.Println("j: ", j, "__start__", start, "__end__", end, " ############ ", "_start:", _start, "_end:", _end, "_err:", err)
		if _err != nil {
			if _err == EOF {
				start = 0
				continue
			} else {
				panic(_err)
			}
		}
		tmparr := values[_start:_end]
		tmparrLength := int64(len(tmparr))
		resultArr[j] = tmparr
		k += tmparrLength
		if k >= pagesize {
			break
		}
		psize = psize - tmparrLength

	}
	return err, resultArr
}

//这里data 必须不是map这种有序的
func BuildDyadicArray(data [][]string) ([][]int, int64) {
	dataLength := len(data)
	arr := make([][]int, dataLength)
	totalNum := int64(0)
	if dataLength == 0 {
		return arr, totalNum
	}

	i := 0
	for _, val := range data {
		valLength := len(val)
		tmparr := make([]int, 0, valLength)
		totalNum += int64(valLength)
		for j, _ := range val {
			tmparr = append(tmparr, j)
		}
		arr[i] = tmparr
		i++
	}
	return arr, totalNum
}
