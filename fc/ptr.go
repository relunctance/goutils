package fc

import "reflect"

//判断是否是指针
func IsPtr(value interface{}) bool {
	v := reflect.ValueOf(value)
	return v.Kind() == reflect.Ptr
}
