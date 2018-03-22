package offsetboundary

import (
	"fmt"
	"utils"
)

//多维数组使用offset和pagesize
//函数功能, 请见测试用例
func GetDyadicArrayByOffset(offset, pagesize, totalNum int64, multiDyadicArr [][]int) (error, [][]int) {
	var err error
	var resultArr [][]int = make([][]int, len(multiDyadicArr))
	if offset > totalNum {
		err = fmt.Errorf("offset:[%d] has exceed DyadicArray TotalLength:[%d]  boundary", offset, totalNum)
	}
	errstr, start, _ := GetBoundary(offset, pagesize, totalNum)
	//fmt.Println("start:", start)
	if errstr != nil {
		err = errstr
		return err, resultArr
	}

	var psize int64 = pagesize
	k := int64(0)
	for j, values := range multiDyadicArr {
		vl := int64(len(values)) //3
		resultArr[j] = nil       //注意: 对应反解析的时候如果碰到为nil说明是需要跳过跳过的 , 参见测试用例: getData()
		if start > vl {
			start = start - vl
			continue
		}

		_err, _start, _end := GetBoundary(start, psize, vl)
		//fmt.Println("_start:", _start)
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
		start = 0 //fixed 重新修正start

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

//针对二维数组获取对应的start和end , 复杂度为O(1)
func GetStartEndByItems(item []int) (start, end int) {
	start = item[0]
	end = item[len(item)-1] + 1
	return
}

/**
 *  indexData 原始 [][]int数组
	offsetIndexData [][]int 经过筛选过后的数组
	//如果结束 返回false
	//没有结束 返回true
*/
func MultiIndexDataHasNext(indexData, offsetIndexData [][]int) bool {

	if len(indexData) != len(offsetIndexData) {
		panic(fmt.Errorf("[indexData] length  not eq [offsetIndexData] length"))
	}

	for key, indexdata := range indexData {
		if len(indexdata) == 0 {
			continue
		}
		if utils.IssetSlice(offsetIndexData, key) {
			item := offsetIndexData[key]
			isend, _ := CheckIsOffsetEnd(indexdata, item)
			if !isend {
				if key != len(indexData)-1 && len(item) == 0 {
					continue
				}
				return true
			}
		}
	}
	return false
}

func CheckIsOffsetEnd(indexData, offsetData []int) (isend bool, next_offset int) {
	iLen := len(indexData)
	oLen := len(offsetData)
	//iLen: 18 oLen: 10
	//iLen: 16 oLen: 0
	if oLen > iLen {
		panic(fmt.Errorf("offsetData length:[%d] has max than indexData length:[%d]\n", oLen, iLen))
	}

	if iLen == 0 {
		return false, 0
	}

	if oLen == 0 {
		return false, 0
	}

	if iLen == oLen {
		return true, 0
	}

	if offsetData[oLen-1] == indexData[iLen-1] {
		return true, 0
	}

	return false, offsetData[oLen-1] + 1
}

//获取二维数组结果
func GetStringSlice(allData [][]string, offsetSlice [][]int) ([][]string, int) {
	res := make([][]string, 0, len(offsetSlice))
	var effectLength int
	for key, items := range offsetSlice {
		if items == nil { //这里如果为nil 说明是跳过的
			continue
		}
		start, end := GetStartEndByItems(items)
		tmpItem := allData[key][start:end]
		effectLength += len(tmpItem)
		res = append(res, tmpItem)

	}
	return res, effectLength

}

//获取一维数组的结果
func GetSingleStringSlice(allData [][]string, offsetSlice [][]int) (res []string, effectLength int) {

	for key, items := range offsetSlice {
		if items == nil { //这里如果为nil 说明是跳过的
			continue
		}
		start, end := GetStartEndByItems(items)
		res = append(res, allData[key][start:end]...)
	}
	effectLength = len(res)
	return
}
