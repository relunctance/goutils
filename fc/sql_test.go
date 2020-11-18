package fc

import (
	"testing"
)

func TestWhere(t *testing.T) {
	sql := "name=? and age =? and cname in(?) and tname in(?)"
	vs := []interface{}{
		"HelloWorld",
		1234,
		[]interface{}{
			"a",
			"b",
			"c",
		},
		[]string{
			"a",
			"b",
			"c",
		},
	}
	where := SqlWhere(sql, vs...)
	expectSql := "name='HelloWorld' and age ='1234' and cname in('a','b','c') and tname in('a','b','c')"
	if where != expectSql {
		t.Fatalf("should be == [%s]", expectSql)
	}
	where = SqlWhere("where "+sql, vs...)
	if where != expectSql {
		t.Fatalf("should be == [%s]", expectSql)
	}
	where = SqlWhere("WHERE "+sql, vs...)
	if where != expectSql {
		t.Fatalf("should be == [%s]", expectSql)
	}
}

func TestSql(t *testing.T) {
	sql := "select * from test where name=? and age =? order by ? limit ?"
	sql = Sql(sql, "helloWorld", 18, "id desc", "500")
	expectSql := "select * from test where name='helloWorld' and age ='18' order by 'id desc' limit '500'"
	if sql != expectSql {
		t.Fatalf("should be == [%s]", expectSql)
	}
}
