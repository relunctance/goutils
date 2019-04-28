package fc

import (
	"strconv"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestIsIp(t *testing.T) {
	ips := map[string]bool{
		`"v=spf1 include:spf.mail.*.*`: false,
	}
	for ip, v := range ips {
		ret := IsIp(ip)
		if ret != v {
			t.Fatalf("should be %v", v)
		}
	}
}
func TestFsSplit(t *testing.T) {
	dataMap := map[string]int{
		"'ipinfo'.*.info.'city'":                      4,
		"'ipinfo'.*.info.city":                        4,
		"ipinfo.*.info.city":                          4,
		"ipinfo.'*'.info.city":                        4,
		"'1234.23.4.2'.ipinfo.'1.0.0.1001'.info.city": 5,
		"'ipinfo'.'1.0.0.1001'.info.name.val.'city'":  6,
		"ipinfo.'1.0.0.1001'.info.city.'a.b.c'":       5,
	}

	for key, val := range dataMap {
		ret := SplitComma(key)
		if len(ret) != val {
			t.Fatalf("[%d] should be == [%d]'", len(ret), val)
		}
	}
}

func TestParseStr(t *testing.T) {
	convey.Convey("测试解析URL", t, func() {
		u := ParseStr("?username=zhangsan&passwd=123456")
		convey.So(u.Get("username"), convey.ShouldEqual, "zhangsan")
		convey.So(u.Get("passwd"), convey.ShouldEqual, "123456")

		convey.So(func() {
			ParseStr("username=zhangsan&passwd=123456")
		}, convey.ShouldPanic)

		convey.So(func() {
			ParseStr("://asdfasdfad?username=zhangsan")
		}, convey.ShouldPanic)
	})
}

func TestOrdAndChr(t *testing.T) {
	convey.Convey("测试ord", t, func() {
		var num rune = 65
		convey.So(Chr(num), convey.ShouldEqual, "A")
	})
	convey.Convey("测试ord", t, func() {
		var str string = "A"
		res := Ord(str)
		convey.So(res[0], convey.ShouldEqual, rune(65))
	})
}

func TestSubStr(t *testing.T) {
	convey.Convey("测试TestSubStr:", t, func() {
		str := "hello world 123456"
		convey.So(Substr(str, 0, 5), convey.ShouldEqual, "hello")
		convey.So("", convey.ShouldEqual, Substr(str, 1000, 5))
		convey.So("hello", convey.ShouldEqual, Substr(str, -1000, 5))
		convey.So("world", convey.ShouldEqual, Substr(str, 6, 5))
	})
}

func TestStrToInt32(t *testing.T) {
	convey.Convey("int32类型转换: ", t, func() {
		i, err := StrToInt32("100")
		convey.So(i, convey.ShouldEqual, 100)
		convey.So(err, convey.ShouldBeNil)

		i, err = StrToInt32("abc")
		convey.So(i, convey.ShouldEqual, 0)
		convey.So(err, convey.ShouldNotBeNil)

		str := strconv.Itoa(1 << 31)
		i, err = StrToInt32(str)
		convey.So(i, convey.ShouldEqual, 0)
		convey.So(err, convey.ShouldNotBeNil)
	})
}

func TestUcFirst(t *testing.T) {
	convey.Convey("测试UcFirst", t, func() {
		convey.So("", convey.ShouldEqual, Ucfirst(""))
		convey.So("Abc", convey.ShouldEqual, Ucfirst("abc"))
		convey.So("A", convey.ShouldEqual, Ucfirst("a"))
		convey.So("Dbc", convey.ShouldEqual, Ucfirst("Dbc"))
		convey.So("1bc", convey.ShouldEqual, Ucfirst("1bc"))
	})
}
