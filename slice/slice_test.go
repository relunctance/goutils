package slice

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestInIntsArray(t *testing.T) {
	//arr := []int{1, 2, 3, 4, 5, 6}
	//InStringArray(3, arr)
	convey.Convey("是否存在Int数组中", t, func() {

		convey.Convey("传入的数组必须是排序后的数组", func() {
			num := 129
			arr := []int{3, 11, 16, 18, 19, 20, 55, 58, 104, 142, 142, 142, 142, 143, 10006, 10015, 10027, 10046, 10053, 10057, 10058, 10059, 10060, 10062, 10063, 10064, 10069, 10221, 129}
			convey.So(func() { InArrayInts(num, arr) }, convey.ShouldPanic)
		})
		convey.Convey("存在于数组中", func() {
			arr := []int{6, 7, 8, 9}
			convey.So(InArrayInts(8, arr), convey.ShouldBeTrue)
		})
		convey.Convey("传入的数组为空", func() {
			arr := make([]int, 0)
			convey.So(InArrayInts(8, arr), convey.ShouldBeFalse)
		})
		convey.Convey("不存在于数组中", func() {
			convey.So(InArrayInts(100, []int{6, 7, 8, 9}), convey.ShouldBeFalse)
		})
	})
}

func TestCheckIntSliceEqual(t *testing.T) {
	convey.Convey("判断int数组是否相等", t, func() {
		convey.Convey("开始检测", func() {
			var b bool
			b = CheckIntSliceEqual([]int{1, 2}, []int{1})
			convey.So(b, convey.ShouldBeFalse)
			b = CheckIntSliceEqual([]int{1, 2}, []int{1, 3})
			convey.So(b, convey.ShouldBeFalse)
			b = CheckIntSliceEqual([]int{1, 2, 18, 20, 88, 9}, []int{1, 2, 18, 20, 88, 9})
			convey.So(b, convey.ShouldBeTrue)
		})
	})
}

func TestCheckStringSliceEqual(t *testing.T) {
	convey.Convey("判断string数组是否相等", t, func() {
		convey.Convey("开始检测", func() {
			var b bool
			b = CheckStringSliceEqual([]string{"a", "b"}, []string{"a"})
			convey.So(b, convey.ShouldBeFalse)
			b = CheckStringSliceEqual([]string{"a", "b"}, []string{"a", "c"})
			convey.So(b, convey.ShouldBeFalse)
			b = CheckStringSliceEqual([]string{"a", "b", "c", "d", "E", "F"}, []string{"a", "b", "c", "d", "E", "F"})
			convey.So(b, convey.ShouldBeTrue)
		})
	})
}

func TestEndStringSlice(t *testing.T) {
	convey.Convey("测试slice获取最后一个值", t, func() {
		convey.Convey("测试字符串", func() {
			arr := []string{
				"a",
				"b",
				"c",
				"d",
				"e",
				"f",
			}
			convey.So("f", convey.ShouldEqual, EndStringSlice(arr))
			convey.So("", convey.ShouldEqual, EndStringSlice([]string{}))
		})

		convey.Convey("测试int数组", func() {
			arr := []int{0, 1, 2, 3, 4, 5, 6}
			convey.So(6, convey.ShouldEqual, EndIntSlice(arr))
			convey.So(-1, convey.ShouldEqual, EndIntSlice([]int{}))
		})

	})
}
