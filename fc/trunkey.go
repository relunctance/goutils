package fc

import (
	"errors"
	"fmt"

	"reflect"
)

var (
	CONTINUE_ERR = errors.New("continue")
)

/*
支持以下4种类型:
	ArrayToSimple([]*User , "Name")
	ArrayToSimple([]User , "Name")
	//需要注意map 是无序的
	ArrayToSimple(map[interface{}]*User , "Name")
	ArrayToSimple(map[interface{}]User , "Name")
*/
func ArrayToSimple(data interface{}, key string) (res []string, err error) {
	res = make([]string, 0, 10)
	var s string
	//使用闭包
	buildData := func(item reflect.Value, key string) (Err error) {
		s, Err = getValueStringByKey(item, key)
		if Err != nil {
			if Err == CONTINUE_ERR {
				res = append(res, "")
				return nil
			}
			return
		}
		res = append(res, s)
		return
	}
	err = commonBuild(data, key, buildData)
	return res, err
}

//根据value 获取对应的string 类型

func getValueStringByKey(item reflect.Value, key string) (vstr string, err error) {
	k := item.Kind()
	if k == reflect.Ptr {
		if item.IsNil() {
			err = CONTINUE_ERR
			return
		}
		item = item.Elem() //转换指针为结构
		k = item.Kind()
	}
	if k != reflect.Struct {
		err = fmt.Errorf("data must is struct like []*struct or []struct")
		return
	}
	_, ok := item.Type().FieldByName(key)
	if !ok {
		err = fmt.Errorf("struct field[%s] not exists", key)
		return
	}
	if item.FieldByName(key).CanInterface() {
		vstr = everything2String(item.FieldByName(key).Interface())
	}
	return
}

func everything2String(v interface{}) string {
	value, ok := v.(string)
	if ok {
		return value
	}
	return fmt.Sprintf("%v", v)

}

/*
DataTrunKey使用示例:

需要注意:
	1. map 是无序的 , 返回结果全部都是无序的
	2. 重复的key 会被覆盖,请尽量保证没有重复的key

示例:
	DataTrunKey([]*User , "Name")
	DataTrunKey([]User , "Name")
	DataTrunKey(map[interface{}]*User , "Name")
	DataTrunKey(map[interface{}]User , "Name")
*/
func DataTrunKey(data interface{}, key string) (res map[string]interface{}, err error) {
	res = make(map[string]interface{})
	var s string
	//使用闭包
	buildData := func(item reflect.Value, key string) (Err error) {
		s, Err = getValueStringByKey(item, key)
		if Err != nil {
			if Err == CONTINUE_ERR {
				return nil
			}
			return
		}
		if item.CanInterface() {
			res[s] = item.Interface()
		}
		return
	}
	err = commonBuild(data, key, buildData)
	return res, err
}

type F func(reflect.Value, string) error

func commonBuild(data interface{}, key string, f F) (err error) {
	v := reflect.ValueOf(data)
	vk := v.Kind()
	if vk != reflect.Slice && vk != reflect.Map {
		err = errors.New("must Slice")
		return
	}

	if v.Len() == 0 {
		return
	}
	l := v.Len()
	switch vk {
	case reflect.Slice:
		for i := 0; i < l; i++ {
			item := v.Index(i)
			err = f(item, key) //闭包
			if err != nil {
				return
			}
		}
	case reflect.Map:
		mapkeys := v.MapKeys()
		for _, idx := range mapkeys {
			item := v.MapIndex(idx)
			err = f(item, key) //闭包
			if err != nil {
				return
			}
		}
	}
	return err
}

func DataTrunMulti(data interface{}, key string) (res map[string][]interface{}, err error) {
	res = make(map[string][]interface{})
	var s string
	//使用闭包
	buildData := func(item reflect.Value, key string) error {
		var Err error
		s, Err = getValueStringByKey(item, key)
		if Err != nil {
			if Err == CONTINUE_ERR {
				return nil
			}
			return Err
		}
		if item.CanInterface() {
			_, ok := res[s]
			if !ok {
				res[s] = make([]interface{}, 0, 5) //减少分配内存
			}
			res[s] = append(res[s], item.Interface())
		}
		return nil
	}
	err = commonBuild(data, key, buildData)
	return
}

//存在覆盖的情况
func Computation(data interface{}, keyField, valueField string) (res map[string]string, err error) {
	res = make(map[string]string)
	var s1, s2 string

	key := keyField + "+" + valueField
	//使用闭包
	buildData := func(item reflect.Value, key string) error {
		field0 := key[0:len(keyField)]
		field1 := key[len(keyField)+1:]
		var Err error
		s1, Err = getValueStringByKey(item, field0)
		if Err != nil {
			if Err == CONTINUE_ERR {
				return nil
			}
			return Err
		}

		s2, Err = getValueStringByKey(item, field1)
		if Err != nil {
			if Err == CONTINUE_ERR {
				return nil
			}
			return Err
		}
		res[s1] = s2
		return nil
	}
	err = commonBuild(data, key, buildData)
	return
}
