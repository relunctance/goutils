package math

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestMinInt(t *testing.T) {
	convey.Convey("测试math:", t, func() {
		convey.Convey("开始测试MinInt():", func() {
			var a, b int = 3, 4
			num := MinInt(a, b)
			convey.So(num, convey.ShouldEqual, a)
			a, b = 18, 3
			num = MinInt(a, b)
			convey.So(num, convey.ShouldEqual, b)
		})
		convey.Convey("开始测试MinInt64():", func() {
			var a, b int64 = 7, 8
			num := MinInt64(a, b)
			convey.So(num, convey.ShouldEqual, a)
			a, b = 18, 3
			num = MinInt64(a, b)
			convey.So(num, convey.ShouldEqual, b)
		})
	})
}
