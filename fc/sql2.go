package fc

import (
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
