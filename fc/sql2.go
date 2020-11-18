package fc

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

// Sql 构建sql
/*
example:
	sql := "select * from test where name=? and age =?  order by ? limit ?"
	vs := []interface{}{
		"hello",
		1234,
		"id desc",
		"1,50",
	}
*/
func Sql(w string, values ...interface{}) string {
	return sqlBuild(w, values...)
}

// 构建where 条件
/*
example:
	sql := "name=? and age =? and cname in(?) and tname in(?)"
	vs := []interface{}{
		"hello",
		1234,
		[]interface{}{ "a", "b", "c" },
		[]string{ "d", "e", "f" },
	}
	SqlWhere(sql, vs...)
*/
func SqlWhere(w string, values ...interface{}) string {
	w = strings.TrimSpace(strings.TrimLeft(strings.TrimLeft(w, "where"), "WHERE")) // 自动移除掉开头 where
	return sqlBuild(w, values...)
}

// 判断传入的元素是否是slice
// 支持检测 []string, []int , []int64 , []in32 , []interface{} , []float64 , []uint32 , []uint64 , ...
func IsSlice(v interface{}) bool {
	return reflect.ValueOf(v).Type().Kind().String() == "slice"
}

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

func ToString(v interface{}) string {
	return fmt.Sprintf("%v", v)
}

// 基于gorm 构建sql
// eg:
//
func gormSqlBuild(sql string, vs ...interface{}) string {
	value := gorm.LogFormatter("sql", "", time.Now().Sub(time.Now()), sql, vs, int64(0))
	return value[3].(string)
}

func sqlBuild(w string, values ...interface{}) string {
	var vt []interface{} = make([]interface{}, 0, len(values))
	for _, v := range values {
		if !IsSlice(v) {
			vt = append(vt, ToString(v))
		} else {
			tmpSlice, _ := SliceToStringSlice(v)
			vt = append(vt, strings.Join(tmpSlice, "','"))
		}
	}
	return gormSqlBuild(w, vt...)
}
