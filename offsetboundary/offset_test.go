package offsetboundary

import (
	"errors"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestCheckIsEndPanic(t *testing.T) {
	convey.Convey("检测CheckIsEnd是否报错", t, func() {
		convey.Convey("必须报错", func() {
			err, _, _ := GetBoundary(11, 5, 10)
			_, err2 := checkIsEnd(err)
			convey.So(err2, convey.ShouldNotBeNil)
		})
	})
}

func TestGetGetHasNext(t *testing.T) {
	convey.Convey("检测GetHasNext:", t, func() {
		convey.Convey("开始检测:", func() {
			var err error = EOF
			hasNext, errRes := GetHasNext(err)
			convey.So(errRes, convey.ShouldBeNil)
			convey.So(hasNext, convey.ShouldBeFalse)
			err = nil
			hasNext, errRes = GetHasNext(err)
			convey.So(hasNext, convey.ShouldBeTrue)
			convey.So(errRes, convey.ShouldBeNil)
			err = errors.New("Faild")
			hasNext, errRes = GetHasNext(err)
			convey.So(hasNext, convey.ShouldBeFalse)
			convey.So(errRes, convey.ShouldNotBeNil)

		})
	})
}

func TestOffset(t *testing.T) {
	convey.Convey("测试GetBoundary Offset", t, func() {
		arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		l := int64(len(arr))
		err, start, end := GetBoundary(5, 5, l)
		convey.So(err, convey.ShouldEqual, EOF)
		convey.So(5, convey.ShouldEqual, start)
		convey.So(10, convey.ShouldEqual, end)

		err, start, end = GetBoundary(10, 10, l)
		convey.So(err, convey.ShouldEqual, EOF)
		convey.So(0, convey.ShouldEqual, start)
		convey.So(0, convey.ShouldEqual, end)

		err, start, end = GetBoundary(9, 1, l)
		convey.So(err, convey.ShouldEqual, EOF)
		convey.So(9, convey.ShouldEqual, start)
		convey.So(10, convey.ShouldEqual, end)

		err, start, end = GetBoundary(9, 3, l)
		convey.So(err, convey.ShouldEqual, EOF)
		convey.So(9, convey.ShouldEqual, start)
		convey.So(10, convey.ShouldEqual, end)

		err, start, end = GetBoundary(6, 3, l)
		convey.So(err, convey.ShouldEqual, nil)
		convey.So(6, convey.ShouldEqual, start)
		convey.So(9, convey.ShouldEqual, end)
	})
}
