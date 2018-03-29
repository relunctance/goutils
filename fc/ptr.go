package fc

import (
	"fmt"
	"reflect"
)

//判断是否是指针
func IsPtr(value interface{}) bool {
	v := reflect.ValueOf(value)
	return v.Kind() == reflect.Ptr
}

// 检测字段是否存在
// 支持[private]字段检测
// 支持[public]字段检测
func FieldExists(fieldName string, ptr interface{}) bool {
	if !IsPtr(ptr) {
		panic(fmt.Errorf("must input a *ptr"))
	}
	v := reflect.ValueOf(ptr)
	e := v.Elem()
	t := e.Type()
	_, ok := t.FieldByName(fieldName)
	return ok
}

// 检测方法是否存在
// 不支持[private]方法检测
// 只支持[public]方法检测,私有方法无法检测出来
func MethodExists(methodName string, ptr interface{}) bool {
	if !IsPtr(ptr) {
		panic(fmt.Errorf("must input a *ptr"))
	}
	typ := reflect.TypeOf(ptr)
	_, ok := typ.MethodByName(methodName)
	return ok
}
