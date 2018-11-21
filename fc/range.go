package fc

import (
	"errors"
)

/*
example:
	RangeByte('a', 'm', 1)
	RangeByte('a', 'z', 2)

*/
func RangeByte(start, end rune, step int) (arr []byte) {
	if step <= 0 {
		panic(errors.New("RangeByte step should not be <=0"))
	}
	iend := int(end)
	istart := int(start)
	if istart >= iend {
		return
	}
	l := (int(end) - int(start)) / step
	arr = make([]byte, l+1)
	j := 0
	for i := istart; i <= iend; i += step {
		arr[j] = byte(i)
		j++

	}
	return

}

/*
example:
	RangeInt(0,10,1)
	RangeInt(0,100,2)
*/
func RangeInt(start, end, step int) (arr []int) {
	if start >= end {
		return
	}
	if step <= 0 {
		panic(errors.New("RangeInt step should not be <=0"))
	}
	l := (end - start) / step
	arr = make([]int, l+1)
	j := 0
	for i := start; i <= end; i += step {
		arr[j] = i
		j++
	}
	return arr
}
