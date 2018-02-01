package slice

import (
	"fmt"
	"sort"
)

func InArrayInts(a int, arr []int) bool {
	if !sort.IntsAreSorted(arr) {
		panic(fmt.Errorf("must be are sorted arr"))
	}
	l := len(arr)
	if l == 0 {
		return false
	}
	key := sort.SearchInts(arr, a) //源码实现是使用二分法实现的, 效率比 for :range 高
	if key > l-1 {                 //肯定不在范围内
		return false
	}
	return arr[key] == a

}

//检测slice是否相等
func CheckIntSliceEqual(x, y []int) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

//检测slice是否相等
func CheckStringSliceEqual(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}
