package fc

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

// 任意[]interface{}转换为[]string
// []string => []string
// []int => []string
// []interface{} => []string
// []int32 => []string
// []uint32 => []string
// []uint64 => []string
// []floag64 => []string
// ...
func SliceToStringSlice(v interface{}) (ret []string, err error) {
	reflectVal := reflect.ValueOf(v)
	if reflectVal.Type().Kind().String() != "slice" {
		return ret, fmt.Errorf("SliceToStringSlice() input value is not a slice ")
	}
	length := reflectVal.Len()
	ret = make([]string, 0, length)
	for i := 0; i < length; i++ {
		rval := reflectVal.Index(i)
		ret = append(ret, ToString(rval.Interface()))
	}
	return ret, nil
}

// 判断传入的元素是否是slice
// 支持检测 []string, []int , []int64 , []in32 , []interface{} , []float64 , []uint32 , []uint64 , ...
func IsSlice(v interface{}) bool {
	return reflect.ValueOf(v).Type().Kind().String() == "slice"
}

func IsArray(value interface{}) bool {
	rv := reflect.ValueOf(value)
	kind := rv.Kind()
	if kind == reflect.Ptr {
		rv = rv.Elem()
		kind = rv.Kind()
	}
	switch kind {
	case reflect.Array, reflect.Slice:
		return true
	default:
		return false
	}
}

// 返回s1中每个元素 可以在s2中的检索到的值
// eg : s1 = ["a" , "b" , "c"] ; s2 = ["a+t" , "b+t"]  => ["a","b"]
func IntersectKeys(s1, s2 []string) (s []string) {
	for _, k1 := range s1 {
		for _, k2 := range s2 {
			if strings.Index(k2, k1) != -1 {
				s = append(s, k1)
			}
		}
	}
	return s
}

// 返回s1中每个元素 无法在s2中的检索到的值
// eg : s1 = ["a" , "b" , "c"] ; s2 = ["a+t" , "b+t"]  => ["c"]
func UniqueKeys(s1, s2 []string) (s []string) {
	for _, k1 := range s1 {
		var has bool = false
		for _, k2 := range s2 {
			if strings.Index(k2, k1) != -1 {
				has = true
				break
			}
		}
		if !has {
			s = append(s, k1)
		}
	}
	return s
}

// diff slice
// Compares arr1 against one or more other arrays and returns the values in arr1 that are not present in any of the other arrays.
func SliceStringDiff(arr1 []string, arrs ...[]string) (data []string) {
	if len(arrs) == 0 {
		return arr1
	}
	i := 0
loop:
	for {
		if i == len(arr1) {
			break
		}
		v := arr1[i]
		for _, arr := range arrs {
			for _, val := range arr {
				if v == val {
					i++
					continue loop
				}
			}
		}
		data = append(data, v)
		i++
	}
	return

}
func SliceChunk(data []string, chunkSize int) (divided [][]string) {

	l := len(data)

	if chunkSize >= l {
		divided = append(divided, data)
		return
	}
	for i := 0; i < l; i += chunkSize {
		end := i + chunkSize
		if end > l {
			end = l
		}

		divided = append(divided, data[i:end])
	}
	return
}

//判断slice中的数组下标是否存在
func IssetSlice(val interface{}, key int) bool {

	var l int

	switch val.(type) {
	case []int:
		v, _ := val.([]int)
		l = len(v)
	case [][]int:
		v, _ := val.([][]int)
		l = len(v)
	case []string:
		v, _ := val.([]string)
		l = len(v)
	case [][]string:
		v, _ := val.([][]string)
		l = len(v)
	case []interface{}:
		v, _ := val.([]interface{})
		l = len(v)
	case [][]interface{}:
		v, _ := val.([][]interface{})
		l = len(v)
	default:
		panic(fmt.Errorf("not support type: %s \n", reflect.TypeOf(val).String()))
	}

	if l > key {
		return true
	}
	return false
}

//元素存在数组检测函数
func InArrayInts(a int, arr []int) bool {
	l := len(arr)
	if l == 0 {
		return false
	}
	if !sort.IntsAreSorted(arr) {
		for _, v := range arr {
			if a == v {
				return true
			}
		}
		return false
	}
	key := sort.SearchInts(arr, a) //源码实现是使用二分法实现的, 效率比 for :range 高
	if key > l-1 {                 //肯定不在范围内
		return false
	}
	return arr[key] == a

}

func InStringArray(v string, arr []string) bool {
	if len(arr) == 0 {
		return false
	}
	for _, val := range arr {
		if v == val {
			return true
		}
	}
	return false
}

//检测slice []int是否相等
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

//检测slice []string是否相等
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

//返回最后一个元素
func EndStringSlice(arr []string) (res string) {
	l := len(arr)
	if l == 0 {
		return
	}
	res = arr[l-1]
	return
}

func EndIntSlice(arr []int) (res int) {
	l := len(arr)
	if l == 0 {
		return -1
	}
	res = arr[l-1]
	return
}

func SliceIntUnique(slice []int) []int {
	sort.Ints(slice)
	i := 0
	var j int
	for {
		if i >= len(slice)-1 {
			break
		}
		for j = i + 1; j < len(slice) && slice[i] == slice[j]; j++ {
		}
		slice = append(slice[:i+1], slice[j:]...)
		i++
	}
	return slice
}

//去重
func SliceStringUnique(slice []string) []string {
	sort.Strings(slice)
	i := 0
	var j int
	for {
		if i >= len(slice)-1 {
			break
		}
		for j = i + 1; j < len(slice) && slice[i] == slice[j]; j++ {
		}
		slice = append(slice[:i+1], slice[j:]...)
		i++
	}
	return slice

}

// 去重+合并
func SliceStringMerge(a []string, b ...[]string) []string {
	return SliceStringUnion(a, b...)
}

// 交集
func SliceStringIntersect(a, b []string) []string {
	ret := make([]string, 0, len(b))
	for _, val := range b {
		if InStringArray(val, a) {
			ret = append(ret, val)
		}
	}
	return ret
}

// 并集
func SliceStringUnion(a []string, b ...[]string) []string {
	for _, v := range b {
		a = append(a, v...)
	}
	return SliceStringUnique(a)
}
